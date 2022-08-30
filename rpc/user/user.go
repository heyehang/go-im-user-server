// Code generated by goctl. DO NOT EDIT!
// Source: user_server.proto

package user

import (
	"context"

	"github.com/heyehang/go-im-grpc/user_server"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AuthUserReply                = user_server.AuthUserReply
	AuthUserReq                  = user_server.AuthUserReq
	CreateUserReply              = user_server.CreateUserReply
	CreateUserReq                = user_server.CreateUserReq
	GetDeviceTokensByUserIDReply = user_server.GetDeviceTokensByUserIDReply
	GetDeviceTokensByUserIDReq   = user_server.GetDeviceTokensByUserIDReq
	IdReq                        = user_server.IdReq
	IdsReq                       = user_server.IdsReq
	ListDeviceToken              = user_server.ListDeviceToken
	LoginReply                   = user_server.LoginReply
	LoginReq                     = user_server.LoginReq
	Token                        = user_server.Token
	UserInfoReply                = user_server.UserInfoReply
	UserInfosReply               = user_server.UserInfosReply

	User interface {
		GetUser(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoReply, error)
		GetUsers(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoReply, error)
		CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserReply, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error)
		AuthUser(ctx context.Context, in *AuthUserReq, opts ...grpc.CallOption) (*AuthUserReply, error)
		GetDeviceTokensByUserID(ctx context.Context, in *GetDeviceTokensByUserIDReq, opts ...grpc.CallOption) (*GetDeviceTokensByUserIDReply, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoReply, error) {
	client := user_server.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) GetUsers(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoReply, error) {
	client := user_server.NewUserClient(m.cli.Conn())
	return client.GetUsers(ctx, in, opts...)
}

func (m *defaultUser) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserReply, error) {
	client := user_server.NewUserClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error) {
	client := user_server.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) AuthUser(ctx context.Context, in *AuthUserReq, opts ...grpc.CallOption) (*AuthUserReply, error) {
	client := user_server.NewUserClient(m.cli.Conn())
	return client.AuthUser(ctx, in, opts...)
}

func (m *defaultUser) GetDeviceTokensByUserID(ctx context.Context, in *GetDeviceTokensByUserIDReq, opts ...grpc.CallOption) (*GetDeviceTokensByUserIDReply, error) {
	client := user_server.NewUserClient(m.cli.Conn())
	return client.GetDeviceTokensByUserID(ctx, in, opts...)
}
