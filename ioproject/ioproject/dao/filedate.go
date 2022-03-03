package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"ioproject/db"
	"ioproject/model"
	"time"
)

// GetLatest 获取最新的存储
func GetLatest(ctx context.Context) (model.FileDate, error) {
	var fileDate model.FileDate
	// 下载 增加条件 下载时间差是0 临时增加
	filter := bson.D{}
	err := db.FileDate.Find(ctx, filter).Sort("-_id").One(&fileDate)
	if err != nil && err.Error() == "mongo: no documents in result" {
		err = nil
	}

	return fileDate, err
}

func PutUpload(ctx context.Context, fileDate *model.FileDate) error {
	_, err := db.FileDate.InsertOne(ctx, fileDate)
	return err
}

// PutDownload 存放下载时间
func PutDownload(ctx context.Context, fileDate *model.FileDate) error {
	update := bson.M{"$set": bson.M{"DownTI": fileDate.DownTI, "EndTime": fileDate.EndTime}}
	err := db.FileDate.UpdateOne(ctx, bson.M{"StartTime": fileDate.StartTime}, update)
	return err
}

// GetEndStorage 根据存储大小读取
func GetEndStorage(ctx context.Context, storage int64) (file model.FileDate, err error) {
	err = db.FileDate.Find(ctx, bson.D{{"EndStorage", storage}}).One(&file)
	return file, err
}

func GetFile(ctx context.Context) (file model.FileDate, err error) {
	err = db.FileDate.Find(ctx, bson.D{{"DownTI", 0}}).Sort("id").One(&file)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return file, nil
	}
	return file, err
}

// GetLatestUnDoLoad 获取没有下载的最新一条数据
func GetLatestUnDoLoad(ctx context.Context) (model.FileDate, error) {
	var fileDate model.FileDate
	filter := bson.D{{"DownTI", 0}}
	err := db.FileDate.Find(ctx, filter).Sort("-_id").One(&fileDate)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return fileDate, nil
	}
	return fileDate, err
}

func Update(ctx context.Context, start, end time.Time) error {
	update := bson.M{"$set": bson.M{"End": end}}
	err := db.FileDate.UpdateOne(ctx, bson.M{"Start": start}, update)
	return err
}
