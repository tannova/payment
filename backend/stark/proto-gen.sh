#!/bin/bash

PROTO_GEN_VALIDATE_LIB="$GOPATH/src/github.com/envoyproxy/protoc-gen-validate"
GOOGLEAPIS_DIR="$GOPATH/src/github.com/googleapis/googleapis"
MONOREPO_DIR="$GOPATH/src/gitlab.com/mcuc/monorepo/backend"
DESCRIPTOR_OUT="api/"

protoc -I. -I"${PROTO_GEN_VALIDATE_LIB}" -I"${GOOGLEAPIS_DIR}" -I"${MONOREPO_DIR}" \
  --include_imports --include_source_info  \
  --descriptor_set_out="stark/api/stark.bin" \
  "stark/api/howard.proto" \
  "stark/api/morgan.proto" \
  "stark/api/pepper.proto" \
  "stark/api/stark.proto" \
  "stark/api/tony.proto" \
  "stark/api/ultron.proto" \
  "stark/api/vision.proto"

