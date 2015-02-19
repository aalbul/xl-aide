package main

import (
	"net/http"
	"github.com/go-martini/martini"
	"strconv"
)

func runServer() {
	m := martini.Classic()
	m.Use(martini.Static("web"))

	m.Get("/import", func(req *http.Request) (int, string) {
			jiraIssue := req.URL.Query()["jiraIssue"][0]
			err := importXlaArchive(jiraIssue)

			if err != nil {
				return 500, err.Error()
			}

			if isParamEnabled(req, "restartServerAfterImport") {
				restartXlDeploy()
			}

			return 200, "XLA attachment has been successfully imported."
		})

	m.Get("/export", func(req *http.Request) {
			jiraIssue := req.URL.Query()["jiraIssue"][0]
			if isParamEnabled(req, "overwriteAlreadyExported") {
				exportXlaArchive(jiraIssue, true)
			}
		})

	m.Run()
}

func isParamEnabled(req *http.Request, paramName string) bool {
	paramValue := req.URL.Query()[paramName]

	if len(paramValue) > 0 {
		enabled,_ := strconv.ParseBool(paramValue[0])
		return enabled;
	}

	return false;
}
