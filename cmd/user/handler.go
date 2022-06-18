package main

import (
	"context"
	"kitexdousheng/kitex_gen/user"
)

// UserSrvImpl implements the last service interface defined in the IDL.
type UserSrvImpl struct{}

// UserRegister implements the UserSrvImpl interface.
func (s *UserSrvImpl) UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// UserLogin implements the UserSrvImpl interface.
func (s *UserSrvImpl) UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserInfo implements the UserSrvImpl interface.
func (s *UserSrvImpl) UserInfo(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}
