package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kitexdousheng/cmd/api/rpc"
	"kitexdousheng/kitex_gen/publish"
	"kitexdousheng/pkg/errno"
	"net/http"
	"path/filepath"
	"time"
)

func PublishAction(c *gin.Context) {
	uid, _ := c.Get("uid")
	userId := uid.(int64)

	data, err := c.FormFile("data")
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	filename := filepath.Base(data.Filename) //返回路径最后的文件名
	finalName := fmt.Sprintf("%d_%d_%s", time.Now().UnixNano(), userId, filename)

	resp, err := rpc.PublishAction(c, &publish.DouyinPublishActionRequest{
		UserId: userId,
		Title:  finalName,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func PublishList(c *gin.Context) {
	uid, _ := c.Get("uid")
	userId := uid.(int64)

	resp, err := rpc.PublishList(c, &publish.DouyinPublishListRequest{
		UserId: userId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	c.JSON(http.StatusOK, resp)
}
