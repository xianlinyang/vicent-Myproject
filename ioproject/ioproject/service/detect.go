package service

import (
	"context"
	"ioproject/dao"
	"ioproject/model"
	"log"
)

/*
	文件验证
	上一次文件是否下载完成
*/

// Detect 文件验证
func Detect(ctx context.Context) (model.FileDate, error) {
	defer log.Println("文件检测完成!")
	// 获取最近完成文件
	file, err := dao.GetLatest(ctx)

	if err != nil {
		return file, err
	}
	// 下载是否完成
	//if file.UpTI != 0 {
	//	if file.DownTI == 0 && file.UpTI != 0 {
	//		// 未完成下载测试
	//		return file, errors.New("file download not completed")
	//	}
	//}

	return file, err
}
