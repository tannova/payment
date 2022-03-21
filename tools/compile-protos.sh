#!/usr/bin/env bash
set -euo pipefail

# Get the directory that the go module is stored in and ensure that it's the correct version.
modpath() {
  set -e
  go mod download "${1}"
  go list -f "{{ .Dir }}" -m "${1}"
}

PROTO_MODULE="."
BACKEND_API="../backend/api"
FRONTEND_API="../frontend/api"
COCOS_API="../cocos/assets/api"

mkdir -p "$BACKEND_API"
mkdir -p "$FRONTEND_API"
mkdir -p "$COCOS_API"

MODULES=$(ls "$PROTO_MODULE")

pg_validate_include_path="$(modpath github.com/envoyproxy/protoc-gen-validate)"

echo "$pg_validate_include_path"



