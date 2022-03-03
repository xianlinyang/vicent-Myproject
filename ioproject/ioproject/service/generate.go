package service

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

var (
	r *rand.Rand
)

func GeneratedMD5ID(file io.Reader) string {
	hash := md5.New()
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	hash.Write(fileByte)
	return hex.EncodeToString(hash.Sum(nil))
}

func DDGenerateFile(size int64) (*File, error) {
	log.Println("文件生成开始")
	file := new(File)
	file.startTime = time.Now()
	// 生成随机数
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成文件
	file1, err := NewRandomManifest(rnd, 1, size)
	if err != nil {
		return nil, err
	}

	buf, err := file1.Archive()
	if err != nil {
		return nil, err
	}
	// 生成id
	file.id = GeneratedMD5ID(bytes.NewReader(buf.Bytes()))
	file.dataReader = bytes.NewReader(buf.Bytes())
	file.name = RandString(17)
	file.size = size

	return file, err
}

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}
