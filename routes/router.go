// @User CPR
package routes

import (
	"VideoWeb/config"
	"VideoWeb/dao"
	"VideoWeb/global"
	"VideoWeb/utils"
	"log"
	"net/http"
	"time"
)

// 初始化全局变量
func InitGlobalVariable() {
	// 初始化Viper
	utils.InitViper()
	// 初始化Logger
	utils.InitLogger()
	// 初始化数据库 DB
	dao.DB = utils.InitMySQLDB()
	// 初始化Redis
	global.Redis = utils.InitRedis()
	// 初始化Casbin
	utils.InitCasbin(dao.DB)
	// 初始化其他参数
	//utils.InitOther()
}

// 后台服务
func BackEndServer() *http.Server {
	backPort := config.Cfg.Server.BackPort
	log.Printf("后台服务启动于 %s 端口", backPort)
	return &http.Server{
		Addr:         backPort,
		Handler:      BackRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

// 前台服务
func FrontEndServer() *http.Server {
	frontPort := config.Cfg.Server.FrontPort
	log.Printf("前台服务启动于 %s 端口", frontPort)
	return &http.Server{
		Addr:         frontPort,
		Handler:      FrontRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
