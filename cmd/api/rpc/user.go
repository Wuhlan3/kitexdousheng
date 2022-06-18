// // Copyright 2021 CloudWeGo Authors
// //
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// //     http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.
// //

package rpc

import (
	"context"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	//etcd "github.com/kitex-contrib/registry-etcd"
	//trace "github.com/kitex-contrib/tracer-opentracing"
	"kitexdousheng/kitex_gen/user"
	"kitexdousheng/kitex_gen/user/usersrv"
	"kitexdousheng/pkg/constants"
	"kitexdousheng/pkg/errno"
	"kitexdousheng/pkg/middleware"
)

var userClient usersrv.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := usersrv.NewClient(
		constants.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func Register(ctx context.Context, req *user.DouyinUserRegisterRequest) error {
	resp, err := userClient.UserRegister(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return nil
}

func Login(ctx context.Context, req *user.DouyinUserLoginRequest) (int64, error) {
	resp, err := userClient.UserLogin(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp.UserId, nil
}
