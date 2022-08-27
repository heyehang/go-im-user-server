// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"go-im-user-server/rpc/internal/logic"
	"go-im-user-server/rpc/internal/svc"
	"go-im-user-server/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IdReq) (*user.UserInfoReply, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

func (s *UserServer) CreateUser(ctx context.Context, in *user.CreateUserReq) (*user.CreateUserReply, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginReply, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) AuthUser(ctx context.Context, in *user.AuthUserReq) (*user.AuthUserReply, error) {
	l := logic.NewAuthUserLogic(ctx, s.svcCtx)
	return l.AuthUser(in)
}
