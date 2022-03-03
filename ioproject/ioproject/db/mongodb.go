package db

import (
	"context"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ioproject/config"

	"log"
)

var (
	client     *qmgo.Client
	FileDate   *qmgo.Collection // qmgo gongju
	SettingDB  *qmgo.Collection
	FeaturesDB *qmgo.Collection
	FileIo     *mongo.Collection
	DataBase   *mongo.Database // go driver
	DemoDB     *qmgo.Collection
)

// Connect 创建mongodb 连接
func Connect(conf config.Config) error {
	option := &qmgo.Config{
		Uri: "mongodb://" + conf.DB_HOST + "/" + conf.DB_NAME,
	}

	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, option)
	if err != nil {
		panic("mongodb connect error" + err.Error())
	}
	FileDate = client.Database(conf.DB_NAME).Collection(conf.DriverTable)
	// todo 配置表
	SettingDB = client.Database(conf.DB_NAME).Collection("config_table")
	// 功能配置 默认false
	FeaturesDB = client.Database(conf.DB_NAME).Collection("features")

	DemoDB = client.Database(conf.DB_NAME).Collection("demo")
	// FileIo = client.Database(conf.DB_NAME).Collection("fileDate")

	newClient, err := mongo.Connect(ctx, options.Client())
	if err != nil {
		return err
	}
	DataBase = newClient.Database(conf.DB_NAME)
	FileIo = newClient.Database(conf.DB_NAME).Collection("fileDate")
	log.Println("mongodb connect success !")
	return err
}
