#!/bin/bash
set -e
# Code coverage generation

COVERAGE_DIR="${COVERAGE_DIR:-coverage}"

# Create the coverage files directory
mkdir -p "$COVERAGE_DIR";

go test -coverprofile="${COVERAGE_DIR}"/coverage.dat ./...

# Display the global code coverage
go tool cover -func="${COVERAGE_DIR}"/coverage.dat ;

if [ "$1" == "html" ]; then
    go tool cover -html="${COVERAGE_DIR}"/coverage.dat -o coverage.html ;
fi
