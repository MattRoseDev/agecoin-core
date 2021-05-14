package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/favecode/agecoin-core/graph/model"
)

func (s *Service) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := s.User.GetUserByUsername(input.Username)
	fmt.Println(err)
	if err != nil {
		return nil, errors.New("username or password is not valid")
	}

	// err = user.ComparePassword(input.Password)
	// if err != nil {
	// 	return nil, errors.New("email/password combination don't work")
	// }

	token, err := user.GenToken()
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (s *Service) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}