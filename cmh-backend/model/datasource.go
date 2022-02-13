package model

import (
	"cmh-backend/cmhtypes"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddSource(userId string, src cmhtypes.Source) bool {
	_, err := DB.Database(userId).Collection("Sources").InsertOne(context.TODO(), src)
	if err != nil {
		log.Default().Println("Error occurred: ", err)
		return false
	} else {
		return true
	}
}

func AddSourceType(userId string, srcType cmhtypes.SourceType) bool {
	//userId used as Db name, as userId will be unique
	_, err := DB.Database(userId).Collection("SourceTypes").InsertOne(context.TODO(), srcType)
	if err != nil {
		log.Default().Println("Error occurred: ", err)
		return false
	} else {
		return true
	}
}

func CheckCollisionInSource(userId string, name string) bool {
	var result bson.M
	err := DB.Database(userId).Collection("Sources").FindOne(context.TODO(), bson.M{"name": name}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		} else {
			log.Default().Println("Error occurred: ", err)
			return true
		}
	} else if len(result) > 0 {
		return true
	}
	return false
}

func CheckCollisionInSourceType(userId string, name string) bool {
	var result bson.M
	err := DB.Database(userId).Collection("SourceTypes").FindOne(context.TODO(), bson.M{"name": name}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		log.Default().Println("Error occurred: ", err)
	} else if len(result) > 0 {
		return true
	}
	return false
}

func FetchList(userId string, listName string) []bson.M {
	var result []bson.M
	var collectionName string

	//Select Collection name
	switch listName {
	case "source":
		collectionName = "Sources"
	case "sourcetype":
		collectionName = "SourceTypes"
	}

	//Extract data from database
	projection := bson.D{{Key: "_id", Value: 0}}
	opts := options.Find().SetProjection(projection)
	cur, err := DB.Database(userId).Collection(collectionName).Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		return []bson.M{{"error": true}}
	}
	err = cur.All(context.TODO(), &result)
	if err != nil {
		return []bson.M{{"error": true}}
	}

	return result
}
