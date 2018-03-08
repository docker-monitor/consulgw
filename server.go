package main

import (
	"io/ioutil"
	"os"

	"github.com/docker-monitor/consulgw/conf"
	"github.com/docker-monitor/consulgw/router"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Init environment
func init() {
	log.Info("Init application configure...")
	var config *conf.Config
	execDirAbsPath, _ := os.Getwd()
	os.Setenv("CONSUL_API", "localhost:8500")
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	data, err := ioutil.ReadFile(execDirAbsPath + "/conf/" + env + ".conf.yaml")
	if err != nil {
		log.Errorln(err)
	}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Errorln(err)
	}

	log.Info("Load environment is ", env)
}

// Main function
func main() {
	router.Start()
}
