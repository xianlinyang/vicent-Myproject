package service

import (
	"context"
	"errors"
	"fmt"
	"ioproject/config"
	"ioproject/dao"
	"ioproject/model"
	"log"
)

var defaultSize int64 = 1000 * 1000 * 1000
var defaultStopSize int64 = 1000 * 1000 * 1000 * 10

func GetSetting(ctx context.Context, cf config.Config) (*model.Setting, error) {
	// 读取数据库配置中的配置
	settings, err := dao.GetSetting(ctx)
	if len(settings) == 0 {
		// 数据库配置不存在,走默认配置
		initSet, err := initSetting(ctx, cf)
		if err != nil {
			return nil, err
		}
		return &initSet, err
	}
	if len(settings) > 1 {
		// todo 目前是根据数据库走,更新配置直接从数据库修改,没有重启以后是根据配置去修改数据库
		return nil, errors.New("存在多个重复配置项")
	}
	if err != nil {
		fmt.Println("===", err)
		return nil, err
	}
	setting := new(model.Setting)
	if len(settings) == 1 {
		setting.FileSize = settings[0].FileSize
		setting.MaxStorage = settings[0].MaxStorage
	}
	return setting, err
}

func initSetting(ctx context.Context, cf config.Config) (model.Setting, error) {
	fileSize, err := config.FormatFileSizeStr(cf.SIZE)
	if err != nil {
		log.Printf("Fail error %v", err)
		fileSize = defaultSize
	}

	stopSize, err := config.FormatFileSizeStr(cf.STOPAPPSIZE)
	if err != nil {
		log.Printf("Fail error %v", err)
		stopSize = defaultStopSize
	}

	setting, err := dao.InitSetting(ctx, model.Setting{
		MaxStorage: stopSize,
		FileSize:   fileSize,
	})
	return setting, err
}
