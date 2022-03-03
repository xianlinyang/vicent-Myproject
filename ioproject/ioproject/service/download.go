package service

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"ioproject/config"
	"ioproject/dao"
	"ioproject/model"
	"log"
	"net/http"
	"sync"
	"time"
)

func Download(ctx context.Context, fileDate *model.FileDate, url string) error {
	var start, end time.Time
	start = time.Now()
	log.Println("开始下载")
	// http get 请求拼装
	data, err := DownLoadGet(url, fileDate.RCID)
	if err != nil {
		return err
	}
	end = time.Now()

	// 下载文件获取MD5
	downloadFileID := GeneratedMD5ID(bytes.NewReader(data))
	// 对比上传下载MD5
	if fileDate.FileID != downloadFileID {
		return errors.New("document content compared error")
	}

	fileDate.DownTI = end.Sub(start).Milliseconds()
	fileDate.EndTime = end
	err = dao.PutDownload(ctx, fileDate)
	if err != nil {
		return err
	}
	log.Println("下载成功")
	return err
}

func DownLoadGet(url, rcid string) ([]byte, error) {
	resp, err := http.Get(url + "/" + rcid)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 响应体
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return data, err
}

/*
	GoDownload 协助下载
	仅仅用于 协程 模型1:1 如果改成多模型需要修改 close(downLoadChan)
*/
func GoDownload(ctx context.Context, conf config.Config, wait *sync.WaitGroup, errChan chan error, downLoadChan, updateTimeChan chan *model.FileDate) {
	defer wait.Done()
	defer close(updateTimeChan)
	for {
		select {
		case file := <-downLoadChan:
			err := Download(ctx, file, conf.URL)
			if err != nil {
				errChan <- err
				log.Println("error", err)
				return
			}
			updateTimeChan <- file
			if file.StartStorage == 0 {
				return
			}
		case <-ctx.Done():
			return
		default:
		}
	}
}

/*
	GoUpDateDownloadTime 协程修改下载文件时间
	仅仅用于 协程 模型1:1 如果改成多模型需要修改 close(downLoadChan)
*/
func GoUpDateDownloadTime(ctx context.Context, wait *sync.WaitGroup, errChan chan error, updateTimeChan chan *model.FileDate) {
	defer wait.Done()
	for {
		select {
		case fileDate := <-updateTimeChan:
			err := dao.PutDownload(ctx, fileDate)
			if err != nil {
				errChan <- err
				log.Println("error", err)
				return
			}
			if fileDate.StartStorage == 0 {
				return
			}
			// fmt.Println("update", fileDate.RCID)
		case <-ctx.Done():
			return
		default:

		}
	}
}
