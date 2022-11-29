package main

import (
	service "golangapp/internal/service"
	"golangapp/pkg"
	"golangapp/pkg/config"
	"golangapp/pkg/logger"
	repositoryInterface "golangapp/pkg/repository"
	repositoryMongoService "golangapp/pkg/repository/mongodb"
)

const (
	configFile = "config.json"
)

func main() {

	pkg.SetWorkingDirectory()

	go Initialize()

	pkg.WaitForTerminate()
}

func Initialize() {

	config.Initialize(configFile)

	logger.CreateLogger(config.GetConfig()["LOG_PREFIX"], config.GetConfig()["LOG_FILENAME"])

	repositoryInterface.SetDB(repositoryMongoService.NewDBEntity(config.GetConfig()["MONGODB_DATABASE"], config.GetConfig()["MONGODB_SERVERURL"]))

	go service.Initialize()

}
