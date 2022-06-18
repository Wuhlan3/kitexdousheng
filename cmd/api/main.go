package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	// "kitexdousheng/config"
	// "kitexdousheng/repository"
	// "kitexdousheng/util"
	"github.com/gin-gonic/gin"
	"kitexdousheng/cmd/api/rpc"
	"net/http"
	//"os"
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

	//r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := http.ListenAndServe(":20000", r); err != nil {
		klog.Fatal(err)
	}
}
