package handlers

import (
	"github.com/gin-gonic/gin"
	"kitexdousheng/cmd/api/rpc"
	"kitexdousheng/kitex_gen/favorite"
	"kitexdousheng/pkg/errno"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	uid, _ := c.Get("uid")
	vid := c.Query("video_id")
	userId := uid.(int64)
	videoId, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	favoriteType := c.Query("action_type")
	temp, err := strconv.ParseInt(favoriteType, 10, 64)
	actionType := int32(temp)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	resp, err := rpc.FavoriteAction(c, &favorite.DouyinFavoriteActionRequest{
		UserId:     userId,
		VideoId:    videoId,
		ActionType: actionType,
	})
	if err != nil || resp == nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	resp, err := rpc.FavoriteList(c, &favorite.DouyinFavoriteListRequest{UserId: userId})
	if err != nil || resp == nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	c.JSON(http.StatusOK, resp)

}
