#!/bin/bash

function build() {
    mm_dir=$(dirname `which $0`)
    if [[ $mm_dir == '.' ]] ; then
        mm_dir=`pwd`
    fi
    export GOPATH=$mm_dir

    go get github.com/twinj/uuid
    go get github.com/acierto/archivex
    go get github.com/acierto/unzipit
    go get github.com/c4pt0r/cfg
    go get github.com/acierto/go-jira-client
    go get launchpad.net/goyaml
    go get github.com/GeertJohan/go.rice
    go get github.com/GeertJohan/go.rice/rice
    go get github.com/stretchr/testify/assert
    go get gopkg.in/yaml.v1

    go build --compiler gccgo
}

build