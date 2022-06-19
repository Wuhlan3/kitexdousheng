package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"kitexdousheng/cmd/feed/repository"
	"kitexdousheng/config"
	feed "kitexdousheng/kitex_gen/feed/feedsrv"
	"kitexdousheng/pkg/constants"
	"kitexdousheng/pkg/middleware"
	"net"
)

func main() {
	//注册ETCD，127.0.0.1:2379
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	//监听的本地ip
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8083")
	if err != nil {
		panic(err)
	}
	//数据库初始化
	repository.Init()
	//config初始化
	config.InitConfig()
	svr := feed.NewServer(new(FeedSrvImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.FeedServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                             // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex，开启连接多路复用
		//server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		//server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler，CPU限流器
		server.WithRegistry(r), // registry，
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
