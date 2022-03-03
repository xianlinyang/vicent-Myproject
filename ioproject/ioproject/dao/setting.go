package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"ioproject/db"
	"ioproject/model"
	"log"
)

func GetSetting(ctx context.Context) ([]*model.Setting, error) {
	var setting []*model.Setting
	err := db.SettingDB.Find(ctx, bson.M{}).All(&setting)
	log.Println("GetSetting", err)
	return setting, err
}

func InitSetting(ctx context.Context, set model.Setting) (model.Setting, error) {
	setting := model.Setting{
		FileSize:   set.FileSize,
		MaxStorage: set.MaxStorage,
	}
	_, err := db.SettingDB.InsertOne(ctx, setting)
	return setting, err
}
