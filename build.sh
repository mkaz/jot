#!/bin/bash

# Build Script for creating multiple architecture releases

# Requires:
# go get github.com/mitchellh/gox

## get version from self to include in file names
go build
VERSION=`jot --version | sed -e 's/jot v//'`

echo "Building $VERSION"
gox -osarch="linux/amd64 darwin/amd64 windows/amd64" -output "{{.Dir}}-$VERSION-{{.OS}}/{{.Dir}}"

for arch in linux darwin windows; do
    tar cf jot-$VERSION-$arch.tar jot-$VERSION-$arch
    gzip jot-$VERSION-$arch.tar
    rm -rf jot-$VERSION-$arch
done

