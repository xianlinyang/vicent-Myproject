package main

import (
	"context"
	"fmt"
	"ioproject/config"
	"ioproject/dao"
	"ioproject/db"
	"ioproject/service"
	"log"
)

func main() {
	ctx := context.Background()
	conf, _ := config.GetConfig()
	err := db.Connect(conf)
	if err != nil {
		log.Println("error ", err)
		return
	}
	err = dao.AddDemo(ctx)
	fmt.Println(err)
	if err != nil {
		return
	}
	// 上传文件
	service.Start(ctx, conf)
	// 顺序下载
	downLoad(ctx, conf)
	// 逆序下载
	// ReverseDownLoad(ctx, conf)
}

func downLoad(ctx context.Context, conf config.Config) {

	// ----------------顺序下载 临时增加----------------
	// 顺序
	for {
		file, err := dao.GetFile(ctx)
		if err != nil {
			return
		}
		if file.UpTI == 0 || file.RCID == "" {
			log.Println("结束测试")
			return
		}
		// 下载
		err = service.Download(ctx, &file, conf.URL)
		if err != nil {
			return
		}
	}
}

func ReverseDownLoad(ctx context.Context, conf config.Config) {

	for {
		file, err := dao.GetLatestUnDoLoad(ctx)
		if err != nil {
			return
		}
		if file.UpTI == 0 || file.RCID == "" {
			log.Println("结束测试")
			return
		}
		// 下载
		err = service.Download(ctx, &file, conf.URL)
		if err != nil {
			return
		}
	}
}
