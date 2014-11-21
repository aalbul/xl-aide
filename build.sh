#!/bin/sh

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

    go build
}

build