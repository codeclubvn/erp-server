package service

import (
	"context"
	"errors"

	"pet-project/model"
	"pet-project/repo"
)

type User struct {
	repo repo.IRepo
}

type IUser interface {
	Login(ctx context.Context, userRequest model.UserRequest) (userResponse model.User, err error)
}

func NewUser(repo repo.IRepo) *User {
	return &User{
		repo: repo,
	}
}

func (s *User) Login(ctx context.Context, userRequest model.UserRequest) (userResponse model.User, err error) {
	user, err := s.repo.GetUserByEmail(userRequest.Username)
	if err != nil {
		return model.User{}, err
	}
	// kiểm tra xem mật khẩu có đúng không
	if user.Password != userRequest.Password {
		return model.User{}, errors.New("Invalid username or password")
	}
	return user, nil
}
