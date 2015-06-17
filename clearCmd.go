package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/bitrise-io/go-pathutil"
	"github.com/codegangsta/cli"
)

func clearCmd(c *cli.Context) {
	isExists, err := pathutil.IsPathExists(currentEnvStoreFilePath)
	if err != nil {
		log.Fatal("Failed to clear EnvStore:", err)
	}
	if !isExists {
		log.Fatal("EnvStore not found in path:", currentEnvStoreFilePath)
	}

	err = writeEnvMapToFile(currentEnvStoreFilePath, []envModel{})
	if err != nil {
		log.Fatal("Failed to clear EnvStore:", err)
	}
	log.Info("EnvStore cleared")
}
