package handlers

import (
	"github.com/gin-gonic/gin"
	"kitexdousheng/cmd/api/rpc"
	"kitexdousheng/kitex_gen/feed"
	"kitexdousheng/pkg/errno"
	"net/http"
)

func Feed(c *gin.Context) {
	myUidInt, _ := c.Get("uid")
	myUid := myUidInt.(int64)
	resp, err := rpc.Feed(c, &feed.DouyinFeedRequest{
		UserId: myUid,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	//视频加载没有出错
	c.JSON(http.StatusOK, resp)
}
