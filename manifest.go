package main

import (
	"log"
	"io/ioutil"
	"fmt"
	"strings"
)

const (
	file_mode = 0644
)

func createManifestForPlugins() string {

	xld_location := getXldLocation()
	dirName := xld_location  + "plugins"
	list, err := ioutil.ReadDir(dirName)

	if err != nil {
		log.Fatalf("ReadDir %s: %v", dirName, err)
	}

	pluginNames := ""
	for _, file := range list {
		if strings.HasSuffix(file.Name(), ".xldp") || strings.HasSuffix(file.Name(), ".jar") {
			pluginNames = fmt.Sprintln(pluginNames + file.Name())
		}
	}

	pluginsMetadata := ".plugins"
	if err := ioutil.WriteFile(pluginsMetadata, []byte(pluginNames), file_mode); err != nil {
		log.Fatalf("WriteFile %s: %v", pluginsMetadata, err)
	}

	return pluginsMetadata
}
