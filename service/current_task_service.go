package service

import (
	"context"
	"errors"

	"github.com/favecode/agecoin-core/graph/model"
	"github.com/favecode/agecoin-core/middleware"
)

func (s *Service) AddCurrentTask(ctx context.Context, input model.AddCurrentTaskInput) (*model.CurrentTask, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	task, _ := s.Task.GetTaskByID(input.TaskID)

	if len(task.ID) < 1 {
		return nil, errors.New("task not found")
	}

	currentTask := &model.CurrentTask{
		TaskID:       input.TaskID,
		UserID:       user.ID,
		DefaultCoins: task.DefaultCoins,
		Active:       input.Active,
	}

	s.CurrentTask.CreateCurrentTask(currentTask)

	return currentTask, nil
}
