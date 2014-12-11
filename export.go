package main

import (
	"github.com/acierto/archivex"
	"log"
	"os"
)

const (
	archive_name = "xla-snapshot"
)


func createArchive() string {

	xld_location := getXldLocation()

	arc := archivex.ZipFile{}

	arc.Create(archive_name)

	createManifestForPlugins()
	arc.AddFile(plugins_metadata)

	for _, dir := range list_of_dirs {
		arc.AddAll(xld_location+dir, true)
	}



	arc.Close()

	return xld_location + full_archive_name
}

func exportXlaArchive(issueId string, replace bool) {
	attachmentPath := createArchive()

	if replace {
		err := jira.UpdateAttachment(issueId, attachmentPath)
		logErrorCleanAndExit(attachmentPath, err)
	} else {
		err := jira.AddAttachment(issueId, attachmentPath)
		logErrorCleanAndExit(attachmentPath, err)
	}

	log.Printf("XLA attachment [%s] has been successfully uploaded.", attachmentPath)
}

func logErrorCleanAndExit(attachmentPath string, err error) {
	os.RemoveAll(attachmentPath)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	os.Remove(plugins_metadata)
}

