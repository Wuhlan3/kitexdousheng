package service

import (
	"kitexdousheng/cmd/repository"
	"kitexdousheng/kitex_gen/user"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
)

func TestUserLogin(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		testName string
		args     args
		wantErr  bool
	}{
		{
			testName: "用户已存在，且密码正确",
			args: args{
				username: "wuhlan3",
				password: "123456",
			},
			wantErr: false,
		},
		{
			testName: "用户已存在，但密码不正确",
			args: args{
				username: "wuhlan3",
				password: "1234567",
			},
			wantErr: true,
		},
		{
			testName: "用户不存在",
			args: args{
				username: "wuhlan66",
				password: "123456",
			},
			wantErr: true,
		},
	}
	repository.Init()
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := UserLogin(&user.DouyinUserLoginRequest{
				Username: tt.args.username,
				Password: tt.args.password,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("UserLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			klog.Info(tt.testName + " success")
		})
	}
}
