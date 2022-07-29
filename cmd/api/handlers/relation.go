package handlers

import (
	"kitexdousheng/cmd/api/rpc"
	"kitexdousheng/kitex_gen/relation"
	"kitexdousheng/pkg/errno"
	"net/http"
	"strconv"

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
	uid, _ := c.Get("uid")
	MyUId := uid.(int64)
	toUserId := c.Query("to_user_id")
	HisUId, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	resp, err := rpc.RelationAction(c, &relation.DouyinRelationActionRequest{
		MyId:       MyUId,
		HisId:      HisUId,
		ActionType: int32(actionType),
	})
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// FollowList 查看登录用户的关注列表
func FollowList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	if userId == 0 {
		uid, _ := c.Get("uid")
		userId = uid.(int64)
	}
	resp, err := rpc.RelationFollowList(c, &relation.DouyinRelationFollowListRequest{UserId: userId})
	if err != nil || resp == nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// FollowerList all users have same follower list 查看粉丝列表
func FollowerList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	if userId == 0 {
		uid, _ := c.Get("uid")
		userId = uid.(int64)
	}
	resp, err := rpc.RelationFollowerList(c, &relation.DouyinRelationFollowerListRequest{UserId: userId})
	if err != nil || resp == nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
