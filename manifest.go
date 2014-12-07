package main

import (
	"log"
	"io/ioutil"
	"fmt"
	"strings"
)

const (
	file_mode = 0644
	plugins_metadata = ".plugins"
)

func readAllArtifacts(dirName string) string {

	pluginNames := ""
	for _, file := range ListDir(dirName) {
		if strings.HasSuffix(file.Name(), ".xldp") || strings.HasSuffix(file.Name(), ".jar") {
			pluginNames = fmt.Sprintln(pluginNames + file.Name())
		}
	}

	return pluginNames
}

func createManifestForPlugins() {

	xld_location := getXldLocation()
	dirName := xld_location  + "plugins"

	pluginNames := readAllArtifacts(dirName)

	if err := ioutil.WriteFile(plugins_metadata, []byte(pluginNames), file_mode); err != nil {
		log.Fatalf("WriteFile %s: %v", plugins_metadata, err)
	}
}
