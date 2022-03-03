package service

import (
	"context"
	"ioproject/config"
	"log"
)

func Start(ctx context.Context, cf config.Config) {
	log.Println("测试")
	for {
		// 读取配置
		setting, err := GetSetting(ctx, cf)
		if err != nil {
			log.Println("配置setting err:", err)
			return
		}
		// 文件检测
		file, err := Detect(ctx)
		if err != nil {
			log.Println("文件检测！ err:", err)
			return
		}
		//if err != nil {
		//	if errors.Is(err, errors.New("file download not completed")) {
		//		errChan <- err
		//		log.Fatalf("Fail error %v", err)
		//		return
		//	} else {
		//		err = Download(ctx, &file, cf.URL)
		//		if err != nil {
		//			errChan <- err
		//			log.Fatalf("Fail error %v", err)
		//			return
		//		}
		//		log.Println("重新下载成功")
		//	}
		//}
		// 储存是否到达指定值
		if file.EndStorage == setting.MaxStorage || (file.StartStorage+setting.FileSize) > setting.MaxStorage {
			log.Println("上传测试完成!")
			return
		}

		// 生成文件
		newFile, err := DDGenerateFile(setting.FileSize)
		if err != nil {
			log.Println("生成文件! err:", err)
			return
		}
		log.Printf("文件生成完成:文件id: %v , 文件大小:%v, 磁盘累计存储:%v", newFile.id, newFile.size, newFile.size+file.EndStorage)
		newFile.size = file.EndStorage
		// 上传文件
		_, err = UploadFile(ctx, cf.URL, newFile, setting.FileSize)
		if err != nil {
			log.Println("err:", err)
			return
		}

	}
}
