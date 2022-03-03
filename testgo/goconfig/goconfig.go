package main


import (
	"github.com/spf13/viper"
	"log"
	"os"
)
func main(){
	//配置文件测试
	//var config Config
	//LoadConfigFromYaml(&config)
	//fmt.Println(config.v.GetString("payment-early"))
}
type Config struct {
	v  *viper.Viper
}
func LoadConfigFromYaml (c *Config) error  {
	c.v = viper.New()

	//设置配置文件的名字
	c.v.SetConfigName("config")

	getwd, _ := os.Getwd()

	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	c.v.AddConfigPath(getwd)
	c.v.AddConfigPath("./")

	//设置配置文件类型
	c.v.SetConfigType("yaml")

	if err := c.v.ReadInConfig(); err != nil{
		return  err
	}

	log.Printf("age: %s, name: %s \n", c.v.Get("information.age"), c.v.Get("information.name"))
	return nil
}