package main

import (
	"log"
	"os"
	"github.com/acierto/unzipit"
)

const (
	full_archive_name = archive_name + ".zip"
)

func importXlaArchive(issueKey string) {
	attachmentPath,_ := jira.DownloadAttachment(issueKey, full_archive_name)

	if attachmentPath == "" {
		log.Printf("Nothing to import. XLA attachment for issue [%s] has not been found.", issueKey)
		os.Exit(1)
	}

	attachment, _ := os.Open(attachmentPath)
	defer attachment.Close()

	unzipit.Unpack(attachment, archive_name)
	os.RemoveAll(attachmentPath)

	for _,dir := range list_of_dirs {
		os.RemoveAll(dir)
		CopyDir(archive_name + string(os.PathSeparator) + dir, dir)
	}

	os.RemoveAll(archive_name)

	log.Print("XLA attachment has been successfully imported.")
}
