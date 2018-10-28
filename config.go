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

type ProducerConf struct {
	Addr       string `yaml:"addr"`
	NotifTopic string `yaml:"notif_topic"`
	MaxRetry int `yaml:"max_retry"`
	IsSuccess bool `yaml:"is_success"`
	RetryBackOff int `yaml:"retry_backoff"`
}

func (pc ProducerConf) GetAddr() string{
	return pc.Addr
}

func (pc ProducerConf) GetTopic()string {
	return pc.NotifTopic
}

func (pc ProducerConf) GetMaxRetry() int{
	return pc.MaxRetry
}

func (pc ProducerConf) GetIsSuccess() bool{
	return pc.IsSuccess
}

func (pc ProducerConf) GetRetryBackoffSec() int{
	return pc.RetryBackOff
}

// Struct for all configs
type Config struct {
	ApiConf      *ApiConf      `yaml:"api"`
	ProducerConf *ProducerConf `yaml:"producer"`
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

func GetProducerConfig() *ProducerConf {
	return config.ProducerConf
}