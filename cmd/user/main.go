package main

import (
	//"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/bound"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	user "kitexdousheng/kitex_gen/user/usersrv"
	"kitexdousheng/pkg/constants"
	"kitexdousheng/pkg/middleware"
	"net"
)

//func main() {
//	svr := user.NewServer(new(UserSrvImpl))
//
//	err := svr.Run()
//
//	if err != nil {
//		log.Println(err.Error())
//	}
//}
func main() {
	//注册ETCD，127.0.0.1:2379
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	//监听的本地ip
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8082")
	if err != nil {
		panic(err)
	}
	//Init()
	svr := user.NewServer(new(UserSrvImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}), // server name
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
