package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"ioproject/db"
	"ioproject/model"
)

func InitFeatures(ctx context.Context, features model.Features) error {
	_, err := db.FeaturesDB.InsertOne(ctx, features)
	return err
}

func GetFeatures(ctx context.Context) ([]model.Features, error) {
	var features []model.Features
	err := db.FeaturesDB.Find(ctx, bson.M{}).All(&features)
	if err != nil && err.Error() == "error mongo: no documents in result" {
		return features, nil
	}
	return features, err
}

func CloseDownLoadFeatures(ctx context.Context, feature model.Features, isClose bool) error {
	update := bson.M{"$set": bson.M{"Download": isClose}}
	err := db.FeaturesDB.UpdateOne(ctx, feature, update)
	//db.FeaturesDB.UpdateOne(ctx, )
	if err != nil && err.Error() == "mongo: no documents in result" {
		return nil
	}
	fmt.Println("CloseDownLoadFeatures", err)
	return err
}

func CloseUpFeatures(ctx context.Context, feature model.Features, isClose bool) error {
	update := bson.M{"$set": bson.M{"UpLoad": isClose}}
	err := db.FeaturesDB.UpdateOne(ctx, feature, update)
	//db.FeaturesDB.UpdateOne(ctx, )
	if err != nil && err.Error() == "mongo: no documents in result" {
		return nil
	}
	return err
}
