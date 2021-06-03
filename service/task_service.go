package service

import (
	"context"
	"errors"

	"github.com/favecode/agecoin-core/graph/model"
	"github.com/favecode/agecoin-core/middleware"
)

func (s *Service) AddTask(ctx context.Context, input model.AddTaskInput) (*model.Task, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	task := &model.Task{
		UserID: user.ID,
		Title: input.Title,
		Description: input.Description,
		DefaultCoins: input.DefaultCoins,
	}

	s.Task.CreateTask(task)
	
	return task, nil
}

func (s *Service) GetTasks(ctx context.Context) ([]*model.Task, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	tasks, err := s.Task.GetTasksByUserId(user.ID)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	
	return tasks, nil
}