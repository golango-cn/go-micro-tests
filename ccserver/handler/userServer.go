package handler

import (
	"cc-go-micro/ccserver/proto/userserver"
	"context"
	"github.com/micro/go-micro/util/log"
)

type UserServer struct {
}

func (u UserServer) Register(ctx context.Context, requ *userserver.RegisterRequ, resp *userserver.RegisterResp) error {

	log.Info("用户注册", requ.UserName, requ.UserEmail)

	resp.UserId = 10000
	return nil

}

func (u UserServer) GetUser(ctx context.Context, requ *userserver.GetUserRequ, resp *userserver.GetUserResp) error {

	resp.UserId = 10000
	resp.UserEmail = "10000@qq.com"
	resp.UserName = "cc"

	return nil

}
