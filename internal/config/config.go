package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	YOUTUBE struct {
		APIKEY string `yaml:"ApiKey"`
	} `yaml:"Youtube"`
	SLACK           string   `yaml:"Slack"`
	TARGET_ACCOUNTS []string `yaml:"TargetAccounts"`
}

var config Config

func init() {
	buf, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetConfig() Config {
	return config
}
