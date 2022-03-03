package main

import (
	"github.com/cihub/seelog"
)

func SetupLogger() {
	logger, err := seelog.LoggerFromConfigAsFile("./seelog/seelog.xml")
	if err != nil {
		return
	}
	seelog.ReplaceLogger(logger)
}

func main() {
	SetupLogger()
	defer seelog.Flush()
	seelog.Infof("我是普通日志")
	seelog.Info("我也是普通日志")
	seelog.Error("我是错误日志")
	seelog.Debug("我是调试日志")
}
