package main

import (
	"flag"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"

	//"mangostreet-ser-iris/app"
	//"mangostreet-ser-iris/common/config"
	//"mangostreet-ser-iris/model"
	"mangostreet-ser-iris/app"
	"mangostreet-ser-iris/common/config"
	"mangostreet-ser-iris/model"
)

var configFile = flag.String("config", "./bbs-go.yaml", "配置文件路径")

func init() {
	flag.Parse()

	config.InitConfig(*configFile)                                                          // 初始化配置
	initLogrus()                                                                            // 初始化日志
	err := simple.OpenDB(config.Conf.MySqlUrl, 5, 20, config.Conf.ShowSql, model.Models...) // 连接数据库
	if err != nil {
		logrus.Error(err)
	}
}

func initLogrus() {
	file, err := os.OpenFile(config.Conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Error(err)
	}
}

func main() {
	app.StartOn()
	app.InitIris()
}
