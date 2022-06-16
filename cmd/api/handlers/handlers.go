package handlers

import (
	"net/http"

	// "github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/errno"
	"kitexdousheng/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
