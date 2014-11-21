package main

import (
	"github.com/c4pt0r/cfg"
	"os"
	"launchpad.net/goyaml"
	"fmt"
	"io/ioutil"
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

func init() {

	xlaConfigPath := GetHomeDir() + string(os.PathSeparator) + xla_config_name

	if ! IsExist(xlaConfigPath) {
		fmt.Printf("Please create xla-config.yml in your home directory. As a template file you can use xla-config-sample.yml")
		os.Exit(1)
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

func ReadXldAppKey(key string) string {
	v, _ := xldConfig.ReadString(key, "")
	return v
}

func GetXlaConfig() *Config {
	return &xlaConfig
}
