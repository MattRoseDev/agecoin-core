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
	}

	s.CurrentTask.CreateCurrentTask(currentTask)

	return currentTask, nil
}

func (s *Service) StartCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	currentTask, _ := s.CurrentTask.GetCurrentTaskByID(currentTaskID)

	if len(currentTask.ID) < 1 {
		return nil, errors.New("current task not found")
	}

	if currentTask.Active == true {
		return nil, errors.New("current task is active")
	}

	s.CurrentTask.DeactiveAllCurrentTaskByUserId(user.ID)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	currentTask.Active = true
	currentTask.Status = 1

	newCurrentTask, err := s.CurrentTask.UpdateCurrentTaskById(currentTask)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	currentTaskHistory := &model.CurrentTaskHistory{
		UserID:        currentTask.UserID,
		CurrentTaskID: currentTask.ID,
		Type:          "START",
	}

	s.CurrentTaskHistory.CreateCurrentTaskHistory(currentTaskHistory)

	return newCurrentTask, nil
}
