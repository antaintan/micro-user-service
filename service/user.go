package service

import (
	"context"
	"huochain/gomicro-demo/user-service/pb"
)

type UserService struct{}

func (s *UserService) Info(ctx context.Context, req *pb.UserRequest, res *pb.UserResponse) error {

	res.Id = req.Id
	res.Name = "bittan 测试"

	return nil
}

func (s *UserService) List(ctx context.Context, req *pb.UserPageRequest, res *pb.UserPageResponse) error {
	return nil
}
