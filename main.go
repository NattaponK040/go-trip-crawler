package main

import (
	"context"
	"go-crawler/client"
	"go-crawler/config"
	"go-crawler/process"
	"go-crawler/repository"
	"os"
	log "github.com/sirupsen/logrus"
)

func main()  {
	log.Info("Start Application")
	conf := config.LoadConfig("", "resource", os.Getenv("ENV"), "application")

	mongoClient, err := client.CreateMongoClient(&conf.Mongoinfo)
	if err != nil {
		log.Error(err)
		os.Exit(0)
	}
	ctx := context.Background()
	data_repos := repository.NewMongoRepository(mongoClient, ctx,&conf.Mongoinfo)
	proc := process.NewProcessor(data_repos)
	proc.Runproc()
}