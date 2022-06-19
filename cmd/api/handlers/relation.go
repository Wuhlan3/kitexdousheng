package handlers

import (
	"github.com/gin-gonic/gin"
)

/*
type Client interface {
	RelationAction(ctx context.Context, Req *relation.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationActionResponse, err error)
	RelationFollowList(ctx context.Context, Req *relation.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationFollowListResponse, err error)
	RelationFollowerList(ctx context.Context, Req *relation.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *relation.DouyinRelationFollowerListResponse, err error)
}
*/

// RelationAction 登录用户给其他用户关注，或者取消关注
func RelationAction(c *gin.Context) {
	//拿到http请求后，参数传给rpc包去调用对应服务返回结果
}

// FollowList 查看登录用户的关注列表
func FollowList(c *gin.Context) {

}

// FollowerList all users have same follower list 查看粉丝列表
func FollowerList(c *gin.Context) {

}
