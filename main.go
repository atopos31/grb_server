package main

import (
	"flag"

	"gvb/core"
	"gvb/global"
	"gvb/routers"
	"gvb/service"
)

func main() {
	// 从命令行读取配置文件位置
	configDir := flag.String("config", "./config/dev.yaml", "config dir")
	flag.Parse()
	// 初始化配置信息
	global.Conf = core.InitConf(*configDir)
	// 初始化日志
	global.Log = core.InitLogger(global.Conf.Logger)

	// 初始化数据库
	db := core.InitGorm(global.Conf.Mysql)
	// 初始化redis
	cache := core.InitRedis(global.Conf.Redis)
	// 初始化站点信息
	siteInfo := core.InitSiteInfo(global.Conf.Sys.SiteInfoPath)
	// service注入
	service.Svc = service.New(db, cache, global.Conf.Oss, global.Conf.Ai, siteInfo)
	// 初始化路由
	router := routers.InitRouter(global.Conf.Sys)
	// 输出配置
	if global.Conf.Sys.Env == "debug" {
		global.Log.Infof("[Config]:%v", global.Conf)
	}

	router.Run(global.Conf.Sys.Addr())
}
