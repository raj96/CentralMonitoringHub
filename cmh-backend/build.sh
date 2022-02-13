#!/bin/bash

BUILDDIR="$(pwd)/build"
EXECNAME="cmh-backend"

if [ ! -d "$BUILDDIR" ]
then
    mkdir "$BUILDDIR"
fi

go build -o "$BUILDDIR/$EXECNAME"
