#!/bin/bash

# Temporary skip
if [ -d report ]; then
    mkdir -p public/report
    mv report/ public/report/
fi

if [ -d frontend/backyard/forge/build ]; then
    mv frontend/backyard/forge/build/ public/
fi