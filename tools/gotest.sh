#!/bin/bash
set -e

go test $(go list ./${NAME}/... | grep -v /vendor/ | grep -v mock) -v -coverprofile cover.out