#!/bin/bash

# This file will run a npm build to create a dist/ directory locally
# and then scp it up to the maverick-studios server.
# afterwards, it'll go into workspace/arcade and git pull to update the files
# and go build on the server and restart the services

# exit when any command fails
set -e

# keep track of the last executed command
trap 'last_command=$current_command; current_command=$BASH_COMMAND' DEBUG
# echo an error message before exiting
trap 'echo "\"${last_command}\" command filed with exit code $?."' EXIT

SITE_HOSTNAME=maverick-studios.ca

pushd frontend
npm run-script build
scp -r dist root@${SITE_HOSTNAME}:~/workspace
popd

ssh root@${SITE_HOSTNAME} << EOF
    cd workspace/arcade;
    git pull origin master
    ./remotecommands.sh
EOF