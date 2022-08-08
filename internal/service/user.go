package service

import (
	"context"
	v1 "realWorld/api/realworld/v1"
)

func (r *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.UserReply, error) {
	// service  ---->   biz
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    "1224853899@qq.com",
			Token:    "",
			Username: "zhiguogg",
			Bio:      "",
			Image:    "",
		}}, nil
}
