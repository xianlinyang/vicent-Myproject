package service

import (
	"context"
	"fmt"
	"ioproject/config"
	"ioproject/dao"
	"ioproject/model"
	"sync"
)

// GetFeatures 获取功能配置
func GetFeatures(ctx context.Context, downloadCh, upLoadCh chan model.Features, closeUpCh, closeDownCh chan struct{}, syWait *sync.WaitGroup) {
	defer syWait.Done()
	defer fmt.Println("已经退出 features")
	err := InitFeatures(ctx)
	var isCloseUpCh, isCloseDownCh bool // isCloseDownLoad
	if err != nil {
		return
	}
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出 feature")
			return
		case <-closeUpCh:
			isCloseUpCh = true
		case <-closeDownCh:
			isCloseDownCh = true
		default:
			fmt.Println("继续")
		}
		fmt.Println("读取配置")
		features, err := dao.GetFeatures(ctx)
		if err != nil {
			return
		}
		isCloseDownCh = features[0].Download
		isCloseUpCh = features[0].UpLoad
		fmt.Println(isCloseDownCh, isCloseUpCh)
		if isCloseUpCh {
			fmt.Println("堵塞1")
			upLoadCh <- features[0]
		}
		fmt.Println("===1")
		//if isCloseDownCh {
		//	fmt.Println("堵塞2")
		//	downloadCh <- features[0]
		//}
	}
}

// DownloadFeatures 从0 开始下载
func DownloadFeatures(ctx context.Context, conf config.Config, downloadCh chan model.Features, errChan chan error, closeDownCh chan struct{}, syWait *sync.WaitGroup) {
	defer syWait.Done()
	defer fmt.Println("===已经退出download")
	var feature model.Features
	for {
		select {
		case isClose := <-downloadCh:
			feature = isClose
		case <-ctx.Done():
			return
		default:

		}
		// fmt.Println("=-----=", feature.Download)
		if feature.Download {
			fmt.Println("收到开始下载")
			file, err := dao.GetFile(ctx)
			if err != nil {
				errChan <- err
				// log.Fatalf("Fail error %v", err)
			}
			if file.UpTI == 0 || file.RCID == "" {
				// 没有上传完成有缓存造成下载时间不准确
				// 关闭下载功能
				// feature 改 false
				err := dao.CloseDownLoadFeatures(ctx, feature, false)
				if err != nil {
					fmt.Println("feature", err)
					errChan <- err
					// log.Fatalf("Fail error %v", err)
				}

				continue
			}
			// 下载
			err = Download(ctx, &file, conf.URL)
			if err != nil {
				errChan <- err
			}
		}
	}
}

func InitFeatures(ctx context.Context) error {
	features, err := dao.GetFeatures(ctx)
	if len(features) != 1 {
		err := dao.InitFeatures(ctx, model.Features{UpLoad: false, Download: false})
		return err
	}

	return err
}
