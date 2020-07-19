#!/bin/bash

if ! command -v godoc &> /dev/null
then
    echo "godoc could not be found, installing godoc"
    go get -u golang.org/x/tools/cmd/godoc
    exit
fi

echo "hosting on port 6060"
echo "visit http://localhost:6060/pkg/github.com/bseto/arcade/backend/"
godoc -http=:6060