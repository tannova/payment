#!/bin/bash
set -e

echo "Running '$0' for '${NAME}'"

docker tag ${NAME}:${CI_COMMIT_SHORT_SHA} $CI_REGISTRY/mcuc/monorepo/${NAME}:${CI_COMMIT_SHORT_SHA}
docker tag ${NAME}:${CI_COMMIT_SHORT_SHA} $CI_REGISTRY/mcuc/monorepo/${NAME}:dev
docker images
docker push $CI_REGISTRY/mcuc/monorepo/${NAME}:${CI_COMMIT_SHORT_SHA}
docker push $CI_REGISTRY/mcuc/monorepo/${NAME}:dev

echo "Deploy application ${NAME}"
export APP_FOLDER="infrastructure/helm/${NAME}/dev"
export API_VERSION=$(grep "appVersion" ${APP_FOLDER}/Chart.yaml | cut -d " " -f2 | cut -d "\"" -f 2)
export RELEASE_NAME="${NAME}"
export DEPLOYS=$(helm ls | grep $RELEASE_NAME | wc -l)
if [ ${DEPLOYS} -eq 0 ]; then helm install ${RELEASE_NAME} ${APP_FOLDER} --namespace=${APP_NAMESPACE}; else helm upgrade ${RELEASE_NAME} ${APP_FOLDER} --namespace=${APP_NAMESPACE}; fi
unset APP_FOLDER
unset API_VERSION
unset RELEASE_NAME
unset DEPLOYS
