package main

import (
	"log"
	"os"
	"github.com/acierto/unzipit"
	"io/ioutil"
	"strings"
	"errors"
	"fmt"
)

const (
	eol = "\n"
	full_archive_name = archive_name + ".zip"
	sep = string(os.PathSeparator)
)

func importXlaArchive(issueKey string) error {
	attachmentPath,err := jira.DownloadAttachment(issueKey, full_archive_name)

	if err != nil {
		return err
	}

	if attachmentPath == "" {
		return errors.New(fmt.Sprintf("Nothing to import. XLA attachment for issue [%s] has not been found.", issueKey))
	}

	attachment, _ := os.Open(attachmentPath)
	defer attachment.Close()

	unpackedFolder := archive_name
	unzipit.Unpack(attachment, unpackedFolder)
	os.RemoveAll(attachmentPath)

	preserverServiceWrapperConfig()

	for _, dir := range list_of_dirs {
		os.RemoveAll(dir)
		CopyDir(unpackedFolder+sep+dir, dir)
	}

	CopyDir(unpackedFolder+sep+"plugins/", "plugins/")

	err = findPluginsDifference(unpackedFolder)

	if err != nil {
		return err;
	}

	os.RemoveAll(unpackedFolder)

	log.Print("XLA attachment has been successfully imported.")

	return nil;
}

func preserverServiceWrapperConfig() {
	dest := archive_name + sep + "conf/"
	confFile := dest + sep + "service-wrapper.conf"
	if !IsExist(dest) {
		os.MkdirAll(dest, 0755)
	}
	if !IsExist(confFile) {
		CopyFile("conf"+sep+"service-wrapper.conf", dest+sep+"service-wrapper.conf")
	}
}

func findPluginsDifference(unpackedFolder string) error {
	pluginsMetadataFile := unpackedFolder + sep + plugins_metadata

	if IsExist(pluginsMetadataFile) {
		content, err := ioutil.ReadFile(pluginsMetadataFile)
		if err != nil {
			log.Fatal(err.Error())
		}

		importedPluginList := strings.Split(string(content), eol)
		foundArtifacts := strings.Split(readAllArtifacts("plugins"), eol)

		diff := difference(importedPluginList, foundArtifacts)
		if cap(diff) > 0 {
			errors.New(fmt.Sprintf("Found the next list of missing plugins: %v. Please install them before proceed further.", diff))
		}
	}

	return nil;
}
