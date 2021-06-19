package biz

import "context"

type User struct {
	Id   string
	Name string
	Age  int
}

type UserRepo interface {
	GetUserById(ctx context.Context, id string) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) GetUserById(ctx context.Context, id string) (*User, error) {
	user, err := uc.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &User{
		Id:   user.Id,
		Name: user.Name,
		Age:  user.Age,
	}, nil
}
