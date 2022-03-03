package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"ioproject/dao"
	"ioproject/model"
	"log"
	"net/http"
	"sync"
	"time"
)

type RCID struct {
	Reference string `json:"reference"`
}

func UploadFile(ctx context.Context, url string, file *File, fileSize int64) (*model.FileDate, error) {

	fileDate := new(model.FileDate)
	var (
		start, end time.Time
	)
	start = time.Now()
	log.Println("开始上传")
	body, err := UpLoadPost(file, url)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}

	// 返回结构体解析
	rcid, err := RCIDJsonToType(body)
	if err != nil {
		log.Println("err1:", err)
		return nil, err
	}
	end = time.Now()

	fileDate.RCID = rcid
	fileDate.FileID = file.id
	fileDate.StartStorage = file.size
	fileDate.EndStorage = fileSize + file.size
	fileDate.UpTI = end.Sub(start).Milliseconds()
	fileDate.StartTime = file.startTime

	err = dao.PutUpload(ctx, fileDate)
	if err != nil {
		log.Println("err3:", err)
		return nil, err
	}
	log.Println("上传成功")
	return fileDate, err
}

func RCIDJsonToType(resp []byte) (string, error) {
	var rcid RCID
	if err := json.Unmarshal(resp, &rcid); err != nil {
		return "", err
	}
	return rcid.Reference, nil
}

func UpLoadPost(file *File, url string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, file.dataReader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Aurora-Index-Document", file.name)
	req.Header.Add("Aurora-Error-Document", file.name)
	req.Header.Add("Aurora-Collection", "false")
	req.Header.Add("Aurora-Collection-Name", file.name)
	req.Header.Add("Aurora-Pin", "false")
	req.Header.Set("Content-Type", "text/plain;charset=utf-8")
	resp, err := (&http.Client{}).Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	// 返回值解析
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	defer func() {
		resp.Body.Close()
	}()

	return body, nil
}

/*
	GoGetFileByEndStorage 通过最后储存值获取文件信息
	协程获取
	GoGetFileByEndStorage *仅仅用于 协程 模型1:1 如果改成多模型需要修改 close(downLoadChan)
*/
func GoGetFile(ctx context.Context, endStorage int64, wait *sync.WaitGroup, errChan chan error, downLoadChan chan *model.FileDate) {
	defer wait.Done()
	defer close(downLoadChan)
	for {
		storage, err := dao.GetFile(ctx)
		if err != nil {
			log.Println("error", err)
			errChan <- err
			return
		}
		if storage.DownTI != 0 {
			continue
		}
		// todo 倒叙读取文件大小 改自定义读取
		if storage.RCID == "" {
			return
			// downLoadChan <- &storage
		}

		// endStorage = storage.StartStorage
		downLoadChan <- &storage
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
