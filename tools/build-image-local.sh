#!/bin/bash

service=$1 
tag=${2:-latest}

build_docker () {
  service=$1
  dockerfile=$2
  IMG="registry.gitlab.com/mcuc/monorepo/$service:$tag"
  docker build -f backend/$dockerfile -t $IMG backend \
    --platform linux/amd64 \
    --build-arg service=$service \
    --build-arg BOT_USER=$BOT_USER \
    --build-arg BOT_PRIVATE_TOKEN=$BOT_PRIVATE_TOKEN \
    --build-arg CI_PROJECT_NAMESPACE=$CI_PROJECT_NAMESPACE
}

build_go () {
  service=$1
  echo "Start build go image $service:$tag"
  build_docker $service go.dockerfile
}

build_envoy () {
  service=$1
  echo "Start build envoy image $service:$tag"
  build_docker $service envoy.dockerfile
}

build_greyhole () {
  service=$1
  version=$2
  echo "Start greyhole service image $service:$tag"
  build_docker $service $service/build/Dockerfile
}

envoy_services="bifrost heimdall"
greyhole_services="shield falcon hawkeye captain nickfury"

regex="$service([^a-zA-Z0-9_-]|$)"
if [[ $envoy_services =~ $regex ]]; then
  build_envoy $service
elif [[ $greyhole_services =~ $regex ]]; then
  build_greyhole $service
else
  build_go $service
fi