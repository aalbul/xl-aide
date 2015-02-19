package main

import (
	"net/http"
	"github.com/go-martini/martini"
	"strconv"
)

func runServer() {
	m := martini.Classic()
	m.Use(martini.Static("web"))

	m.Get("/import", func(req *http.Request) {
			jiraIssue := req.URL.Query()["jiraIssue"][0]
			restartServerAfterImport := req.URL.Query()["restartServerAfterImport"][0]
			importXlaArchive(jiraIssue)
			enabled,_ := strconv.ParseBool(restartServerAfterImport)
			if enabled {
				restartXlDeploy()
			}
		})

	m.Get("/export", func(req *http.Request) {
			jiraIssue := req.URL.Query()["jiraIssue"][0]
			overwriteAlreadyExported := req.URL.Query()["overwriteAlreadyExported"][0]
			enabled,_ := strconv.ParseBool(overwriteAlreadyExported)
			if enabled {
				exportXlaArchive(jiraIssue, true)
			}
		})

	m.Run()
}
