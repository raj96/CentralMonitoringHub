package model

import (
	"cmh-backend/cmhtypes"
	"context"
	"errors"
	"fmt"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Internal calls only
func FetchSourceAndTypes() map[string]cmhtypes.Statistics {
	rawDbList, err := DB.ListDatabaseNames(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	excludeNames := []string{"admin", "local"}
	var dbList []string

	for _, name := range rawDbList {
		if sort.SearchStrings(excludeNames, name) > 1 {
			dbList = append(dbList, name)
		}
	}

	allSrcTypes := make(map[string]cmhtypes.Statistics)
	projection := bson.D{{Key: "_id", Value: 0}}
	opts := options.Find().SetProjection(projection)

	var srcTypeResult cmhtypes.SourceType
	for _, dbName := range dbList {
		cursor, err := DB.Database(dbName).Collection("SourceTypes").Find(context.TODO(), bson.M{}, opts)
		if err != nil {
			panic(err)
		}
		for cursor.TryNext(context.TODO()) {
			err := cursor.Decode(&srcTypeResult)
			if err != nil {
				continue
			}
			allSrcTypes[srcTypeResult.Name] = srcTypeResult.Stats
		}
	}

	var srcResult cmhtypes.Source
	srcMap := make(map[string]cmhtypes.Statistics)
	for _, dbName := range dbList {
		cursor, err := DB.Database(dbName).Collection("Sources").Find(context.TODO(), bson.M{}, opts)
		if err != nil {
			panic(err)
		}

		for cursor.TryNext(context.TODO()) {
			cursor.Decode(&srcResult)
			srcMap[srcResult.Name] = allSrcTypes[srcResult.SourceTypeName]
		}
	}

	return srcMap
}

func FetchSpecificSourceType(userId string, sourceName string) (cmhtypes.Statistics, error) {
	var srcResult cmhtypes.Source
	var ctxt context.Context
	ctxt, cancelCtxt := context.WithCancel(ctxt)
	var err error

	err = DB.Database(userId).Collection("Sources").FindOne(ctxt, bson.M{"name": sourceName}).Decode(&srcResult)
	if ctxt.Err() != nil {
		cancelCtxt()
		return nil, ctxt.Err()
	}
	if err != nil {
		cancelCtxt()
		return nil, err
	}
	srcTypeName := srcResult.SourceTypeName

	var srcTypeResult cmhtypes.SourceType
	err = DB.Database(userId).Collection("SourceTypes").FindOne(ctxt, bson.M{"name": srcTypeName}).Decode(&srcTypeResult)
	if ctxt.Err() != nil {
		cancelCtxt()
		return nil, ctxt.Err()
	}
	if err != nil {
		cancelCtxt()
		return nil, err
	}

	cancelCtxt()
	return srcTypeResult.Stats, nil
}

func CreateTimeSeries(name string) bool {
	opts := options.CreateCollection()
	opts.SetTimeSeriesOptions(options.TimeSeries().SetGranularity("seconds").SetTimeField("timestamp").SetMetaField("metadata"))

	err := DB.Database("Statistics").CreateCollection(context.TODO(), name, opts)
	fmt.Println(err)
	if err == nil {
		StatisticsCollectionNames[name] = true
	}
	return err == nil
}

func InsertDataToTimeSeries(name string, data map[string]interface{}) error {
	if !StatisticsCollectionNames[name] {
		return errors.New(name + " not created")
	}
	_, err := DB.Database("Statistics").Collection(name).InsertOne(context.TODO(), data)
	return err
}
