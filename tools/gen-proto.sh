#!/bin/sh

service=$1

docker run -it \
-e BOT_USER=$BOT_USER \
-e BOT_PRIVATE_TOKEN=$BOT_PRIVATE_TOKEN \
-e CI_PROJECT_NAMESPACE=$CI_PROJECT_NAMESPACE \
-v $PROJECT_DIR:/go/src/gitlab.com/$CI_PROJECT_NAMESPACE/monorepo \
registry.gitlab.com/$CI_PROJECT_NAMESPACE/monorepo/publish-proto:latest /bin/bash \
/go/src/gitlab.com/$CI_PROJECT_NAMESPACE/monorepo/tools/night-kit.sh $service
