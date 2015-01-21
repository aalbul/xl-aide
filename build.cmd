@echo off

REM
REM Batch script to build the XL Aide
REM

set GOPATH=%~dp0

go get gopkg.in/yaml.v1
go get github.com/twinj/uuid
go get github.com/acierto/archivex
go get github.com/acierto/unzipit
go get github.com/c4pt0r/cfg
go get github.com/acierto/go-jira-client
go get launchpad.net/goyaml
go get github.com/GeertJohan/go.rice
go get github.com/GeertJohan/go.rice/rice
go get github.com/stretchr/testify/assert

go build
