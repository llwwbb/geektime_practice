package service

import (
	"context"
	v1 "github.com/llwwbb/geektime_practice/go_advanced_training/week4/api/user/service/v1"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc *biz.UserUseCase
}

var _ v1.UserServer = &UserService{}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserRes, error) {
	user, err := s.uc.GetUserById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetUserRes{
		Id:   user.Id,
		Name: user.Name,
		Age:  int32(user.Age),
	}, nil
}
