package model

import (
	b1 "gopkg.in/mgo.v2/bson"
	"time"
)

type FileDate struct {
	FileID       string    `bson:"FileID"`       // 文件id
	UpTI         int64     `bson:"UpTI"`         // 上传时间差
	DownTI       int64     `bson:"DownTI"`       // 下载时间差
	RCID         string    `bson:"RCID"`         // 返回的文件id
	StartTime    time.Time `bson:"StartTime"`    // 开始时间
	EndTime      time.Time `bson:"EndTime"`      // 下载时间
	StartStorage int64     `bson:"StartStorage"` // 开始文件累计大小
	EndStorage   int64     `bson:"EndStorage"`   // 结束文件累计大小
	//RAWAIT    float64   `bson:"RAWAIT"`
	//WAWAIT    float64   `bson:"WAWAIT"`
	//Util      float64   `bson:"Util"`
	//MAXRAWAIT float64   `bson:"MAXRAWAIT"`
	//MAXWAWAIT float64   `bson:"MAXWAWAIT"`
}

type Setting struct {
	// 生成文件配置
	FileSize int64 `bson:"FileSize"`
	// 累计文件大小配置
	MaxStorage int64 `bson:"MaxStorage"`
}

// Features 功能
type Features struct {
	ID b1.ObjectId `bson:"_id"`
	// 上传文件默认false
	UpLoad bool `bson:"UpLoad"`
	// 下载文件默认false
	Download bool `bson:"Download"`
}

type FileModel struct {
	ID    string
	Start time.Time
	File  []byte
	End   time.Time
}

type Demo struct {
	Name string `bson:"FileID"`
	// 上传文件默认false
}
