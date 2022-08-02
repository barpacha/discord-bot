package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Main struct {
	Discord Discord `yaml:"discord"`
}

type Discord struct {
	Token string `yaml:"token"`
}

func Yaml(path string) Main {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("read config error: %s", err)
	}
	var configStruct Main
	err = yaml.Unmarshal(file, &configStruct)
	if err != nil {
		log.Fatalf("unmarshal config file error: %s", err)
	}
	return configStruct
}
