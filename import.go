package main

import (
	"log"
	"os"
	"github.com/acierto/unzipit"
	"io/ioutil"
	"strings"
)

const (
	eol = "\n"
	full_archive_name = archive_name + ".zip"
	sep = string(os.PathSeparator)
)

func importXlaArchive(issueKey string) {
	attachmentPath,err := jira.DownloadAttachment(issueKey, full_archive_name)

	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}

	if attachmentPath == "" {
		log.Printf("Nothing to import. XLA attachment for issue [%s] has not been found.", issueKey)
		os.Exit(1)
	}

	attachment, _ := os.Open(attachmentPath)
	defer attachment.Close()

	unpackedFolder := archive_name
	unzipit.Unpack(attachment, unpackedFolder)
	os.RemoveAll(attachmentPath)

	for _,dir := range list_of_dirs {
		os.RemoveAll(dir)
		CopyDir(unpackedFolder + sep + dir, dir)
	}

	findPluginsDifference(unpackedFolder)
	os.RemoveAll(unpackedFolder)

	log.Print("XLA attachment has been successfully imported.")
}

func findPluginsDifference(unpackedFolder string) {
	pluginsMetadataFile := unpackedFolder + sep + plugins_metadata

	content, err := ioutil.ReadFile(pluginsMetadataFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	importedPluginList := strings.Split(string(content), eol)
	foundArtifacts := strings.Split(readAllArtifacts(unpackedFolder), eol)

	log.Printf("Found the next list of missing plugins: %v. Please install them before proceed further.", difference(importedPluginList, foundArtifacts))
	os.Exit(1)
}
