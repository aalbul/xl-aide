package main

import (
	"github.com/c4pt0r/cfg"
	"os"
	"launchpad.net/goyaml"
	"fmt"
	"io/ioutil"
	"github.com/GeertJohan/go.rice"
	"log"
	"gopkg.in/acierto/yaml.v1"
)

var xldConfig *cfg.Cfg
var xlaConfig = Config{}

type Config struct {
	Jira 	Jira  `yaml:"jira"`
	Xld 	Xld  `yaml:"xld"`
}

type Xld struct {
	Login        string `yaml:"login"`
	Password     string `yaml:"password"`
}

type Jira struct {
	Host         string `yaml:"host"`
	ApiPath      string `yaml:"api_path"`
	ActivityPath string `yaml:"activity_path"`
	Login        string `yaml:"login"`
	Password     string `yaml:"password"`
}

const (
	xla_config_name = "xla-config.yml"
)

var xlaConfigPath = GetHomeDir() + string(os.PathSeparator) + xla_config_name

func init() {

	if ! IsExist(xlaConfigPath) {
		conf := rice.Config{
	LocateOrder: []rice.LocateMethod{rice.LocateEmbedded, rice.LocateAppended, rice.LocateFS},
	}
	box, err := conf.FindBox("conf")

	configSampleFileContent, err := box.String("xla-config-sample.yml")
	if err != nil {
		log.Fatal(err)
	}

	WriteToFile(xlaConfigPath, []byte(configSampleFileContent))
}

file, e := ioutil.ReadFile(xlaConfigPath)
if e != nil {
fmt.Printf("Config file error: %v\n", e)
os.Exit(1)
}

sep := string(os.PathSeparator)
xldConfig = cfg.NewCfg("conf" + sep + "deployit.conf")
xldConfig.Load()

err := goyaml.Unmarshal([]byte(file), &xlaConfig)
if err != nil {
panic(err)
}
}

func ValidateConfig() {

	if xlaConfig.Jira.Login == "" {
		log.Fatalf("Please provide your Jira login in [%s] file.", xlaConfigPath)
		os.Exit(1)
	}

	if xlaConfig.Jira.Password == "" {
		log.Fatalf("Please provide your Jira password in [%s] file.", xlaConfigPath)
		os.Exit(1)
	} else if !isBase64(xlaConfig.Jira.Password) {
		xlaConfig.Jira.Password = encode(xlaConfig.Jira.Password)
	}

	if xlaConfig.Xld.Login == "" {
		log.Fatalf("Please provide your XLD admin login in [%s] file.", xlaConfigPath)
		os.Exit(1)
	}

	if xlaConfig.Xld.Password == "" {
		log.Fatalf("Please provide your XLD admin password in [%s] file.", xlaConfigPath)
		os.Exit(1)
	} else if !isBase64(xlaConfig.Xld.Password) {
		xlaConfig.Xld.Password = encode(xlaConfig.Xld.Password)
	}

	d, _ := yaml.Marshal(&xlaConfig)

	err := ioutil.WriteFile(xlaConfigPath, d, 0644)
	if err != nil {
		log.Fatalf("During updating of xla config file happened an error: %s ", err.Error())
		os.Exit(1)
	}
}

func ReadXldAppKey(key string) string {
	v, _ := xldConfig.ReadString(key, "")
	return v
}

func GetXlaConfig() *Config {
	return &xlaConfig
}
