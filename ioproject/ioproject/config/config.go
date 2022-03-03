package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var transformSize int64 = 1000

type Config struct {
	DB_HOST      string `yaml:"DB_HOST"`
	DB_USER      string `yaml:"DB_USER"`
	DB_PWD       string `yaml:"DB_PWD"`
	DB_NAME      string `yaml:"DB_NAME"`
	TOKENKEY     string `yaml:"TOKENKEY"`
	APPPORT      string `yaml:"APPPORT"`
	SIZE         string `yaml:"FILE_SIZE"`
	URL          string `yaml:"URL"`
	STOPAPPTIME  string `yaml:"STOP_APP_TIME"`
	STOPAPPSIZE  string `yaml:"STOP_APP_SIZE"`
	Disk         string `yaml:"DISK"`
	DriverTable  string `yaml:"DRIVER_TABLE"`
	DownLoadSize string `yaml:"DOWN_LOAD_SIZE"`
	Time         string `yaml:"Time"`
	//	ConfigTable  string `yaml:"Config_Table"`
}

func GetConfig() (Config, error) {
	root, err := os.Getwd()
	if err != nil {
		log.Fatalf("Fail error %v", err)
		return Config{}, err
	}
	filePath := root + "/config.yaml"
	buffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Fail error %v", err)
		return Config{}, err
	}
	config := Config{}
	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		log.Fatalf("Fail error %v", err)
		return Config{}, err
	}
	return config, err
}

func FormatFileSizeStr(fileSize string) (int64, error) {
	switch fileSize != "" {
	case strings.Contains(strings.ToUpper(fileSize), "MB") || strings.Contains(strings.ToUpper(fileSize), "M"):
		fileSize = StringTrimSuffix(fileSize, "MB", "M")
		fs, err := strconv.ParseInt(fileSize, 10, 64)
		return fs * transformSize * transformSize, err
	case strings.Contains(strings.ToUpper(fileSize), "TB") || strings.Contains(strings.ToUpper(fileSize), "T"):
		fileSize = StringTrimSuffix(fileSize, "tb", "T")
		fs, err := strconv.ParseInt(fileSize, 10, 64)
		return fs * transformSize * transformSize * transformSize * transformSize, err
	case strings.Contains(strings.ToUpper(fileSize), "KB") || strings.Contains(strings.ToUpper(fileSize), "K"):
		fileSize = StringTrimSuffix(fileSize, "kb", "K")
		fs, err := strconv.ParseInt(fileSize, 10, 64)
		return fs * transformSize, err
	case strings.Contains(strings.ToUpper(fileSize), "GB") || strings.Contains(strings.ToUpper(fileSize), "G"):
		fileSize = StringTrimSuffix(fileSize, "GB", "G")
		fs, err := strconv.ParseInt(fileSize, 10, 64)
		return fs * transformSize * transformSize * transformSize, err
	default:
		fileSize = StringTrimSuffix(fileSize, "B", "b")
		fs, err := strconv.ParseInt(fileSize, 10, 64)
		return fs * transformSize * transformSize * transformSize, err
	}
}

func StringTrimSuffix(old, up, lower string) string {
	old = strings.TrimSuffix(strings.ToUpper(old), up)
	old = strings.TrimSuffix(strings.ToUpper(old), lower)

	return old
}
