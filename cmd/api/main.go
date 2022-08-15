package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"kitexdousheng/cmd/api/rpc"
	"kitexdousheng/config"
	"kitexdousheng/pkg/constants"
	"kitexdousheng/pkg/tracer"
	"net/http"
)


func main() {

	Init()
	r := gin.Default()

	initRouter(r)

	config.InitConfig()


	//r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := http.ListenAndServe(":"+viper.GetString("server.port"), r); err != nil {
		klog.Fatal(err)
	}
}

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

// kLogInit 重定向klog的输出
//func kLogInit(path string){
//	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//	klog.SetOutput(f)
//}