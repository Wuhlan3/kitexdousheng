package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"kitexdousheng/kitex_gen/relation"
	"kitexdousheng/kitex_gen/relation/relationsrv"
	"kitexdousheng/pkg/constants"
	"kitexdousheng/pkg/errno"
	"kitexdousheng/pkg/middleware"
	"log"
	"time"
)

/*
对这三个函数做一下封装就好了，供api模块的handler去调用
type Client interface {
	RelationAction(ctx context.Context, Req *relation.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationActionResponse, err error)
	RelationFollowList(ctx context.Context, Req *relation.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationFollowListResponse, err error)
	RelationFollowerList(ctx context.Context, Req *relation.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationFollowerListResponse, err error)
}
*/

var relationClient relationsrv.Client

func initRelationRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := relationsrv.NewClient(
		constants.RelationServiceName,
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
	relationClient = c
}
func RelationAction(ctx context.Context, Req *relation.DouyinRelationActionRequest) (*relation.DouyinRelationActionResponse, error) {
	resp, err := relationClient.RelationAction(ctx, Req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
func RelationFollowList(ctx context.Context, Req *relation.DouyinRelationFollowListRequest) (*relation.DouyinRelationFollowListResponse, error) {
	resp, err := relationClient.RelationFollowList(ctx, Req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
func RelationFollowerList(ctx context.Context, Req *relation.DouyinRelationFollowerListRequest) (*relation.DouyinRelationFollowerListResponse, error) {
	resp, err := relationClient.RelationFollowerList(ctx, Req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
