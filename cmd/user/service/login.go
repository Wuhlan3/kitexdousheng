package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"kitexdousheng/cmd/repository/db"
	"kitexdousheng/kitex_gen/user"
)

type UserLoginFlow struct {
	name     string
	password string
}

func UserLogin(req *user.DouyinUserLoginRequest) (int64, error) {
	return NewUserLoginFlow(req.Username, req.Password).Do()
}

func NewUserLoginFlow(name string, password string) *UserLoginFlow {
	return &UserLoginFlow{
		name:     name,
		password: password,
	}
}

func (f *UserLoginFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	id, err := f.login()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (f *UserLoginFlow) checkParam() error {
	return nil
}

func (f *UserLoginFlow) login() (int64, error) {
	tempUser, err := db.NewUserDaoInstance().QueryUserByName(f.name)
	if err != nil || tempUser == nil {
		return 0, errors.New("用户不存在")
	}
	err = bcrypt.CompareHashAndPassword([]byte(tempUser.Password), []byte(f.password))
	if err != nil {
		return 0, errors.New("密码错误")
	}

	return tempUser.Id, nil
}
