#!/bin/bash

mm_dir=$(dirname `which $0`)
if [[ $mm_dir == '.' ]] ; then
    mm_dir=`pwd`
fi
export GOPATH=$mm_dir

./build.sh

go test .