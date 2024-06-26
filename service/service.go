package service

import (
	"context"
	"fmt"

	"github.com/favecode/agecoin-core/database"
	"github.com/favecode/agecoin-core/graph/model"
)

type Service struct {
	User               database.User
	Password           database.Password
	Task               database.Task
	TaskHistory database.TaskHistory
}

func New(service Service) *Service {
	return &Service{
		User:               service.User,
		Password:           service.Password,
		Task:        service.Task,
		TaskHistory: service.TaskHistory,
	}
}

func (s *Service) Test(ctx context.Context, input model.TestInput) (*model.Test, error) {
	panic(fmt.Errorf("not implemented"))
}
