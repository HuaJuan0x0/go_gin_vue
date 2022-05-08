package main

import (
	"go_gin_vue/common"
	_ "go_gin_vue/controller"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	common.Setup()
	r := gin.Default()
	r = CollectRoute(r)
	if port := viper.GetString("server.port"); port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
