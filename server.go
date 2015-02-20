package main

import (
	"net/http"
	"github.com/go-martini/martini"
	"strconv"
	"encoding/json"
	"github.com/acierto/go-jira-client"
)

func runServer() {
	m := martini.Classic()

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

	m.Get("/export", func(req *http.Request) (int, string) {
			jiraIssue := req.URL.Query()["jiraIssue"][0]

			err := exportXlaArchive(jiraIssue, isParamEnabled(req, "overwriteAlreadyExported"))

			if err != nil {
				return 500, err.Error()
			}

			return 200, "XLA attachment has been successfully uploaded."
		})

	m.Get("/jiraHost", func(req *http.Request) (int, string) {
			return 200, GetXlaConfig().Jira.Host
		})

	m.Get("/pick", func(req *http.Request) (int, string) {
			request := gojira.IssuePickRequest {
				Query: req.URL.Query()["query"][0],
				ShowSubTasks: isParamEnabled(req, "showSubTasks"),
				ShowSubTaskParent: isParamEnabled(req, "showSubTaskParent"),
			}
			issues := jira.PickIssues(&request)
	
			jsonString, _ := json.Marshal(issues)
			return 200, string(jsonString)
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
