package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Main struct {
	Discord Discord `yaml:"discord"`
	Store   Store   `yaml:"store"`
}

type Discord struct {
	Token string `yaml:"token"`
}

type Store struct {
	Database               string `yaml:"database"`
	MaxOpenConns           int    `yaml:"max_open_conns"`
	MaxIdleConns           int    `yaml:"max_idle_conns"`
	ConnMaxLifetimeMinutes int    `yaml:"conn_max_lifetime_minutes"`
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
