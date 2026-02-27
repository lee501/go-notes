package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type SC struct {
	Database
}

type Database struct {
	Host   string `mapstructure:"host"`
	User   string `mapstructure:"user"`
	Dbname string `mapstructure:"dbname"`
	Pwd    string `mapstructure:"pwd"`
	On     int    `mapstructure:"id"`
}

func viperDemo() {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := viper.New()

	config.AddConfigPath(path)     //设置读取的文件路径
	config.SetConfigName("config") //设置读取的文件名
	config.SetConfigType("yaml")   //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	var sc SC
	config.Unmarshal(&sc)
	fmt.Println(sc.On)
	//打印文件读取出来的内容:
	//fmt.Println(config.Get("database.host"))
	//fmt.Println(config.Get("database.user"))
	//fmt.Println(config.Get("database.dbname"))
	//fmt.Println(config.Get("database.pwd")

}
