package dao

import (
	"context"
	"ioproject/db"
	"ioproject/model"
)

func AddDemo(ctx context.Context) error {
	var features model.Demo
	features.Name = "zhangsan"
	_, err := db.DemoDB.InsertOne(ctx, features)

	return err
}
