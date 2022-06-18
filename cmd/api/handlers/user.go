package handlers

import (
	"context"
	"kitexdousheng/cmd/api/rpc"
	"kitexdousheng/kitex_gen/user"
	"kitexdousheng/pkg/errno"
	//"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	//数据解析
	username := c.Query("username")
	password := c.Query("password")
	//数据验证
	if len(username) == 0 || len(password) < 5 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	//rpc通信
	err := rpc.Register(context.Background(), &user.DouyinUserRegisterRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func Login(c *gin.Context) {
	//数据解析
	username := c.Query("username")
	password := c.Query("password")
	//数据验证
	if len(username) == 0 || len(password) < 5 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	//rpc通信
	userId, err := rpc.Login(context.Background(), &user.DouyinUserLoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, userId)
}
