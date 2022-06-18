package handlers

import (
	"context"
	"kitexdousheng/cmd/api/rpc"
	"kitexdousheng/kitex_gen/user"
	"kitexdousheng/pkg/errno"
	"kitexdousheng/pkg/middleware"
	"net/http"
	"strconv"

	//"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterResponse struct {
	*user.DouyinUserRegisterResponse
	Token string `json:"token"`
}
type LoginResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserId     int64  `json:"user_id,omitempty"`
	Token      string `json:"token"`
}
type UserInfoResponse struct {
	*user.DouyinUserResponse
}

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
	registerResponse, err := rpc.Register(context.Background(), &user.DouyinUserRegisterRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	token, err := middleware.ReleaseToken(registerResponse.UserId)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	c.JSON(http.StatusOK, RegisterResponse{registerResponse, token})
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
	token, err := middleware.ReleaseToken(userId)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	c.JSON(http.StatusOK, LoginResponse{
		StatusCode: 0,
		StatusMsg:  "",
		UserId:     userId,
		Token:      token,
	})
}

func UserInfo(c *gin.Context) {
	//数据解析
	myUidInt, _ := c.Get("uid")
	myUid := myUidInt.(int64)
	id := c.Query("user_id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	resp, err := rpc.UserInfo(c, &user.DouyinUserRequest{
		MyId:  myUid,
		HisId: userId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	c.JSON(http.StatusOK, resp)

}
