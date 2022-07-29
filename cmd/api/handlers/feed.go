package handlers

import (
	"kitexdousheng/cmd/api/rpc"
	"kitexdousheng/kitex_gen/feed"
	"kitexdousheng/pkg/errno"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	// 如果没有传latest_time，则默认为当前时间*******************************
	var CurrentTimeInt = time.Now().UnixMilli()
	var CurrentTime = strconv.FormatInt(CurrentTimeInt, 10)
	var LatestTimeStr = c.DefaultQuery("latest_time", CurrentTime)
	LatestTime, err := strconv.ParseInt(LatestTimeStr, 10, 64)
	if err != nil {
		// 无法解析latest_time
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	//**********************************************************************

	myUidInt, _ := c.Get("uid")
	myUid := myUidInt.(int64)

	resp, err := rpc.Feed(c, &feed.DouyinFeedRequest{
		UserId:     myUid,
		LatestTime: &LatestTime,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	//视频加载没有出错
	c.JSON(http.StatusOK, resp)
}
