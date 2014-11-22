package main

import (
	"github.com/acierto/go-jira-client"
)

func GetJira() *gojira.Jira {

	config := GetXlaConfig().Jira

	return gojira.NewJira(
		config.Host,
		config.ApiPath,
		config.ActivityPath,
		&gojira.Auth{config.Login, config.Password,},
	)

}
