package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"kitexdousheng/kitex_gen/favorite"
	"kitexdousheng/kitex_gen/favorite/favoritesrv"
	"kitexdousheng/pkg/constants"
	"kitexdousheng/pkg/errno"
	"kitexdousheng/pkg/middleware"
	"time"
)

var favoriteClient favoritesrv.Client

func initFavoriteRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := favoritesrv.NewClient(
		constants.FavoriteServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}
func FavoriteAction(ctx context.Context, Req *favorite.DouyinFavoriteActionRequest) (*favorite.DouyinFavoriteActionResponse, error) {
	resp, err := favoriteClient.FavoriteAction(ctx, Req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
func FavoriteList(ctx context.Context, Req *favorite.DouyinFavoriteListRequest) (*favorite.DouyinFavoriteListResponse, error) {
	resp, err := favoriteClient.FavoriteList(ctx, Req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
