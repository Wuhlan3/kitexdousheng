package main

import (
	"context"
	"kitexdousheng/cmd/user/service"
	"kitexdousheng/kitex_gen/user"
	"kitexdousheng/pkg/errno"
)

// UserSrvImpl implements the last service interface defined in the IDL.
type UserSrvImpl struct{}

// UserRegister implements the UserSrvImpl interface.
func (s *UserSrvImpl) UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	id, err := service.UserRegister(req.Username, req.Password)
	if err != nil {
		return &user.DouyinUserRegisterResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  &errno.ServiceErr.ErrMsg,
		}, err
	}
	return &user.DouyinUserRegisterResponse{
		UserId:     id,
		StatusCode: errno.Success.ErrCode,
		StatusMsg:  &errno.Success.ErrMsg,
	}, nil
}

// UserLogin implements the UserSrvImpl interface.
func (s *UserSrvImpl) UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	userId, err := service.UserLogin(req)
	if err != nil {
		return &user.DouyinUserLoginResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  &errno.ServiceErr.ErrMsg,
		}, err
	}
	return &user.DouyinUserLoginResponse{
		UserId:     userId,
		StatusCode: errno.Success.ErrCode,
		StatusMsg:  &errno.Success.ErrMsg,
	}, nil
}

// UserInfo implements the UserSrvImpl interface.
func (s *UserSrvImpl) UserInfo(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	User, err := service.QueryUserInfo(req.MyId, req.HisId)
	if err != nil {
		return nil, err
	}
	return &user.DouyinUserResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		User: &user.User{
			Id:            User.Id,
			Name:          User.Name,
			FollowCount:   &User.FollowCount,
			FollowerCount: &User.FollowerCount,
			IsFollow:      false,
		},
	}, nil
}
