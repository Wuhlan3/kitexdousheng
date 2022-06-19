package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"kitexdousheng/cmd/api/rpc"
	"kitexdousheng/config"
	"net/http"
)

func main() {
	rpc.InitRPC()
	// config.InitConfig() // 设置配置文件
	// if err := repository.Init(); err != nil {
	// 	os.Exit(-1)
	// } //数据库连接
	// util.InitLogger()
	r := gin.Default()

	initRouter(r)

	config.InitConfig()

	//r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := http.ListenAndServe(":"+viper.GetString("server.port"), r); err != nil {
		klog.Fatal(err)
	}
}
