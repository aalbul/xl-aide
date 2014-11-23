package main

import (
	"github.com/acierto/archivex"
	"log"
)

const (
	archive_name = "xla-snapshot"
)


func createArchive() string {

	xld_location := getXldLocation()

	arc := archivex.ZipFile{}

	arc.Create(archive_name)

	manifest := createManifestForPlugins()
	arc.AddFile(manifest)

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

