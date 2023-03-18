package main

import (
	"flag"
	"fmt"

	"go_api_skeleton/bootstrap"
	btsConfig "go_api_skeleton/config"
	"go_api_skeleton/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Logger
	bootstrap.SetupLogger()

	// 设置 gin 的运行模式，支持 debug, release, test
	gin.SetMode(gin.ReleaseMode)

	// new 一个 Gin Engine 实例
	router := gin.New()
	// 初始化 DB
	bootstrap.SetupDB()
	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
