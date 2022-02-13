package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var StatisticsCollectionNames map[string]bool

func ConnectToDB(connectionString string) {
	clientOptions := options.Client().ApplyURI(connectionString)
	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err == nil {
		log.Default().Println("Connected To DB")
		StatisticsCollectionNames = make(map[string]bool)
		DB = db
		statsCollNames, err := DB.Database("Statistics").ListCollectionNames(context.TODO(), map[string]string{})
		if err != nil {
			panic(err)
		}

		for _, name := range statsCollNames {
			StatisticsCollectionNames[name] = true
		}
	} else {
		fmt.Println(err)
		panic(err)
	}
}

func DisconnectFromDB() {
	DB.Disconnect(context.TODO())
}
