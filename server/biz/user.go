package biz

import (
	"context"
	v1 "github.com/heyujiang/hapis/protogen-go/common/v1"
	userV1 "github.com/heyujiang/hapis/protogen-go/user/v1"
	"github.com/heyujiang/user/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type UserBiz struct {
	conf config.UserBiz
	userV1.UnimplementedUserServer
}

func StartUserBiz(conf config.UserBiz) {
	bizServ := &UserBiz{
		conf: conf,
	}

	listener, err := net.Listen("tcp", bizServ.conf.Url)
	if err != nil {
		panic(err)
	}

	grpcSer := grpc.NewServer()
	userV1.RegisterUserServer(grpcSer, bizServ)
	reflection.Register(grpcSer)

	if err := grpcSer.Serve(listener); err != nil {
		panic(err)
	}
}

func (u *UserBiz) Login(ctx context.Context, req *userV1.UserLoginReq) (*v1.Result, error) {
	return &v1.Result{Code: v1.StatusCode_Success, Msg: "Login success"}, nil
}

func (u *UserBiz) GetUserInfo(ctx context.Context, req *userV1.GetUserInfoReq) (*userV1.GetUserInfoResp, error) {
	return &userV1.GetUserInfoResp{}, nil
}
