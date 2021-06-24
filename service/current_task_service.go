package service

import (
	"context"
	"errors"

	"github.com/favecode/agecoin-core/graph/model"
	"github.com/favecode/agecoin-core/middleware"
	"github.com/favecode/agecoin-core/util"
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

	if currentTask.Active == bool(true) {
		return nil, errors.New("current task is active")
	}

	s.CurrentTask.DeactiveAllCurrentTaskByUserId(user.ID)

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

func (s *Service) PauseCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	currentTask, _ := s.CurrentTask.GetCurrentTaskByID(currentTaskID)

	if len(currentTask.ID) < 1 {
		return nil, errors.New("current task not found")
	}

	if currentTask.Active == bool(false) {
		return nil, errors.New("current task is not active")
	}

	currentTaskHistory, _ := s.CurrentTaskHistory.GetCurrentTaskHistoryByCurrentTaskId(currentTaskID)

	coins := util.CalculateCurrentTaskCoins(currentTaskHistory.CreatedAt)

	currentTask.Active = bool(false)
	currentTask.Status = 1
	currentTask.Coins = &coins

	newCurrentTask, err := s.CurrentTask.UpdateCurrentTaskById(currentTask)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	s.CurrentTask.DeactiveAllCurrentTaskByUserId(user.ID)

	newCurrentTaskHistory := &model.CurrentTaskHistory{
		UserID:        currentTask.UserID,
		CurrentTaskID: currentTask.ID,
		Type:          "PAUSE",
	}

	s.CurrentTaskHistory.CreateCurrentTaskHistory(newCurrentTaskHistory)

	return newCurrentTask, nil
}

func (s *Service) DeleteCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	currentTask, _ := s.CurrentTask.GetCurrentTaskByID(currentTaskID)

	if len(currentTask.ID) < 1 {
		return nil, errors.New("current task not found")
	}

	if currentTask.UserID != user.ID {
		return nil, errors.New("authorization failed")
	}

	deletedCurrentTask, err := s.CurrentTask.DeleteCurrentTaskById(currentTaskID)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	newCurrentTaskHistory := &model.CurrentTaskHistory{
		UserID:        currentTask.UserID,
		CurrentTaskID: currentTask.ID,
		Type:          "CANCEL",
	}

	s.CurrentTaskHistory.CreateCurrentTaskHistory(newCurrentTaskHistory)

	return deletedCurrentTask, nil
}
