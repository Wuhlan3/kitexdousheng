package handlers

import (
	"github.com/gin-gonic/gin"
	"kitexdousheng/cmd/api/rpc"
	"kitexdousheng/kitex_gen/comment"
	"kitexdousheng/pkg/errno"

	"net/http"
	"strconv"
)

func CommentAction(c *gin.Context) {
	uid, _ := c.Get("uid")
	userId := uid.(int64)

	vid := c.Query("video_id")
	VideoId, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	actionType := c.Query("action_type")

	if actionType == "1" {
		text := c.Query("comment_text")
		//err := rpc.CommentAction(userId, VideoId, text)
		resp, err := rpc.CommentAction(c, &comment.DouyinCommentActionRequest{
			UserId:      userId,
			VideoId:     VideoId,
			CommentText: &text,
		})
		if err != nil {
			SendResponse(c, errno.ConvertErr(err), nil)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
	return
}

func CommentList(c *gin.Context) {
	vid := c.Query("video_id")
	VideoId, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	resp, err := rpc.CommentList(c, &comment.DouyinCommentListRequest{
		VideoId: VideoId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	c.JSON(http.StatusOK, resp)
}
