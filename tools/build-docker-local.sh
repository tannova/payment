#!/bin/bash
set -euo pipefail

# Remember that we stand at root directory to run this script

find backend -name 'Makefile' | grep -v 'vendor' | grep -v 'pkg' | grep -v 'backend/Makefile' | while read item; do
    echo "build '${item}'"
    # Main build command.
    INCLUDE_MAKEFILE=$item make release-local
    STATUS=FAILED
    
    if [ $? -eq 0 ]; then
        STATUS=SUCCESS
    fi
    echo "build '${item}' done with status '${STATUS}'"	
done

echo "clean up build stage image..."
docker image prune -f
rm -f Dockerfile
