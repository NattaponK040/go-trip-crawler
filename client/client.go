package client

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-crawler/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoClient(cf *config.Mongodb) (*mongo.Client, error) {
	option := options.Client().ApplyURI(cf.Url).SetMinPoolSize(cf.MaxPool).SetMinPoolSize(cf.MinPool)
	mongoClient, err := mongo.Connect(context.TODO(), option)
	if err != nil {
		return nil, err
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Error(err)
	}
	log.Info("MongoDB Connected!")
	return mongoClient, nil
}
