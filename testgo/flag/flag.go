package main

import (
	"flag"
	"fmt"
)

func main(){
	data_path := flag.String("config","","DB data path")
	flag.Parse()
	fmt.Println(data_path)
	//c := getcommand()
	//globalFlags := c.root.PersistentFlags()
	//globalFlags.StringVar(&c.globalConfigFile, "config", "", "config file (default is $HOME/.aurorakeeper.yaml)")
	//
	//c.root.Execute()
}
