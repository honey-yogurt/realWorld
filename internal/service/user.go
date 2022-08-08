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

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.UserReply, err error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    u.Email,
			Username: u.Username,
			Token:    u.Token,
		},
	}, nil
}
