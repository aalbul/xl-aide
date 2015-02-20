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

	arc.AddWithExcludedExtensions(xld_location+"plugins/", []string{".jar", ".xldp"})

	arc.Close()

	return xld_location + full_archive_name
}

func exportXlaArchive(issueId string, replace bool) error {
	attachmentPath := createArchive()

	var err error

	if replace {
		err = jira.UpdateAttachment(issueId, attachmentPath)
		clean(attachmentPath)
	} else {
		err = jira.AddAttachment(issueId, attachmentPath)
		clean(attachmentPath)
	}

	log.Printf("XLA attachment [%s] has been successfully uploaded.", attachmentPath)

	return err;
}

func clean(attachmentPath string) {
	os.RemoveAll(attachmentPath)
	os.Remove(plugins_metadata)
}

