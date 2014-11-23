@echo off

REM
REM Batch script to build the XL Aide
REM

set GOPATH=%~dp0

go get github.com/twinj/uuid
go get github.com/acierto/archivex
go get github.com/acierto/unzipit
go get github.com/c4pt0r/cfg
go get github.com/acierto/go-jira-client
go get launchpad.net/goyaml

go build
