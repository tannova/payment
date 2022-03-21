#!/usr/bin/env bash

docker --version; rc=$?
if [ $rc -ne 0 ]; then
  # Somehow, this is important even though we already have `setup_remote_docker`.
  VER="18.06.1-ce"
  curl -L -o /tmp/docker-$VER.tgz https://download.docker.com/linux/static/stable/x86_64/docker-$VER.tgz
  tar -xz -C /tmp -f /tmp/docker-$VER.tgz
  mv /tmp/docker/* /usr/bin
  docker --version
fi

# Extract commit range (or single commit).
if [ "${CI_COMMIT_BEFORE_SHA}" = "0000000000000000000000000000000000000000" ];
then
  COMMIT_RANGE="${CI_MERGE_REQUEST_DIFF_BASE_SHA}..."
else
  COMMIT_RANGE="${CI_COMMIT_BEFORE_SHA} ${CI_COMMIT_SHA}"
fi

# This is the list of all makefiles that we've already built. We don't include the 
# root makefile by default.
TESTED=`readlink -e ${PWD}/Makefile`
echo "${TESTED}" > testedlist

# Main test function. Takes a directory as input.
test () {
  echo "Build input = $1"
  DIRNAME=`echo $1`
  MKFILE=`echo "${DIRNAME}/Makefile"`
  SLASHES=${PWD//[^\/]/}

  # Try walking up the path until we find a makefile.
  for (( n=${#SLASHES}; n>0; --n )); do
    if [ -f $MKFILE ]; then
      echo "Found Makefile in ${DIRNAME}"
      break
    else
      DIRNAME="${DIRNAME}/.."
      MKFILE=`echo "${DIRNAME}/Makefile"`
    fi
  done
  
  # Get the full path of the makefile.
  MKFILE_FULL=`readlink -e ${MKFILE}`
  
  # Build only if it's not on our list of built makefiles.
  TESTED=$(<testedlist)
  if [[ $TESTED != *"${MKFILE_FULL}"* ]]; then
    echo "Build ${DIRNAME} (${MKFILE_FULL})"
    
    # Main test command.
    INCLUDE_MAKEFILE=$MKFILE make test


    ### Modify accordingly.
    STATUS=FAILED
    ERROR=$?
    if [ $? -eq 0 ]; then
      STATUS=SUCCESS
    fi

    USER_EMAIL=$GITLAB_USER_EMAIL
    USERNAME=${USER_EMAIL%@*}
    
    SLACK_URL=`echo "https://hooks.slack.com/services/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"`
    PAYLOAD=`printf '{"text":"\`#%s\` %s\n\`\`\`\nTestInput: %s\nMakefile: %s\nBranch: %s\nCommitRange: %s\nStatus: %s\n\`\`\`"}' "$CI_PIPELINE_URL" @"$USERNAME" "$DIRNAME" "$MKFILE_FULL" "$CI_MERGE_REQUEST_SOURCE_PROJECT_URL" "$COMMIT_RANGE" "$STATUS"`
    curl -X POST -d "payload=${PAYLOAD}" $SLACK_URL
    
    # Add item to our list of built makefiles.
    TESTED=`echo "${TESTED};${MKFILE_FULL}"`
    echo "${TESTED}" > testedlist
  else
    echo "Skip ${MKFILE_FULL} (already built, or root)"
  fi
}

# Pretest function. Takes a file as input.
processline () {
  line=$1
  echo "processline ${line}"

  if [[ $line == vendor* ]] || [[ $line == pkg* ]]; then
    # The changed line is common. We will iterate through all dirs except hidden ones, 'vendor',
    # and 'pkg' to see if test is necessary.
    find . -type d -not -path "*/\.*" | grep -v 'vendor' | grep -v 'pkg' | while read item; do
      # Get the current package's full list of golang dependencies (recursive).
      PKG_GODEPS=`go list -f '{{ .Deps }}' $item`
      
      if [ $? -eq 0 ]; then
        LINE_DIR=`dirname $line`

        # See if this package has a dependency with the changed file. If so, proceed with test.
        if [[ $PKG_GODEPS = *"${LINE_DIR}"* ]]; then
          echo "'${item}' has a dependency with '${LINE_DIR}'"
          # Remove the './' prefix (output from 'find' command).
          TO_TEST=`echo "${item}" | cut -c 3-`
          test $TO_TEST
        fi
      fi
    done
  else
    # The changed line belongs to either a service or cmd.
    TO_TEST=`dirname $line`
    test $TO_TEST
  fi
}

echo "Commit range ${COMMIT_RANGE}"


if [[ $COMMIT_RANGE = *" "* ]]; then
  # Unfortunately we don't always get a commit range from circleci.
  # Walk through each changed file within the commit.
  echo "No commit range? (${COMMIT_RANGE})"
  git diff-tree --no-commit-id --name-only -r $COMMIT_RANGE | while read line; do
    processline $line
    echo "-"
  done
else
  # Walk through each changed file within the commit range.
  echo "Proper commit range = ${COMMIT_RANGE}"
  git diff --name-only $COMMIT_RANGE | while read line; do
    processline $line
    echo "-"
  done
fi
