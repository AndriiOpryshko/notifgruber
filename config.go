package main

import (
	"path/filepath"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// Config for connection to btc wallet
type ApiConf struct {
	Addr string `yaml:"addr"`
	Port int `yaml:"port"`
}

type KafkaConf struct {
	Addr string `yaml:"addr"`
	NotifTopic string `yaml:"notif_topic"`
}

// Struct for all configs
type Config struct {
	ApiConf ApiConf `yaml:"api"`
	KafkaConf KafkaConf `yaml:"kafka"`
}

// Config global object
var config *Config = &Config{}



// Init all configs from config.yml
func InitConfig() {
	absPath, _ := filepath.Abs("./config.yml")

	log.WithFields(log.Fields{
		"path": absPath,
	}).Info("Start inition config")

	yamlFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.WithFields(log.Fields{
			"path": absPath,
			"err": err,
		}).Fatal("Initialization config error")
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.WithFields(log.Fields{
			"path": absPath,
			"err": err,
		}).Fatal("Initialization config error")
	}

	log.WithFields(log.Fields{
		"success": true,
	}).Info("Inition config")
}
