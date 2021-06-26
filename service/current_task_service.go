package service

import (
	"context"
	"errors"

	"github.com/favecode/agecoin-core/graph/model"
	"github.com/favecode/agecoin-core/middleware"
	"github.com/favecode/agecoin-core/util"
)

func (s *Service) saveCurrentTaskHistory(currentTask *model.CurrentTask, historyType string) {
	currentTaskHistory := &model.CurrentTaskHistory{
		UserID:        currentTask.UserID,
		CurrentTaskID: currentTask.ID,
		Type:          historyType,
	}

	s.CurrentTaskHistory.CreateCurrentTaskHistory(currentTaskHistory)
}

func (s *Service) getCurrentTaskCoins(currentTaskID string) int {
	currentTaskHistory, _ := s.CurrentTaskHistory.GetCurrentTaskHistoryByCurrentTaskId(currentTaskID)

	if len(currentTaskHistory.ID) < 1 {
		return 0
	} else {
		return util.CalculateCurrentTaskCoins(currentTaskHistory.CreatedAt)
	}
}

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

	if currentTask.Status == 2 {
		return nil, errors.New("current task is finished")
	}

	s.CurrentTask.DeactiveAllCurrentTaskByUserId(user.ID)

	currentTask.Active = true
	currentTask.Status = 1

	newCurrentTask, err := s.CurrentTask.UpdateCurrentTaskById(currentTask)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	s.saveCurrentTaskHistory(currentTask, "START")

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

	s.saveCurrentTaskHistory(currentTask, "PAUSE")

	return newCurrentTask, nil
}

func (s *Service) FinishCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	currentTask, _ := s.CurrentTask.GetCurrentTaskByID(currentTaskID)

	if len(currentTask.ID) < 1 {
		return nil, errors.New("current task not found")
	}

	if currentTask.Status == 2 {
		return nil, errors.New("current task is finished")
	}

	if currentTask.Active == bool(true) {
		coins := s.getCurrentTaskCoins(currentTask.ID)
		currentTask.Active = bool(false)
		currentTask.Coins = &coins
	}

	currentTask.Status = 2

	newCurrentTask, err := s.CurrentTask.UpdateCurrentTaskById(currentTask)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	s.CurrentTask.DeactiveAllCurrentTaskByUserId(user.ID)

	s.saveCurrentTaskHistory(currentTask, "FINISH")

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

	s.saveCurrentTaskHistory(currentTask, "CANCEL")

	return deletedCurrentTask, nil
}

func (s *Service) GetCurrentTasks(ctx context.Context) ([]*model.CurrentTask, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	activeCurrentTask, _ := s.CurrentTask.GetActiveCurrentTaskByUserId(user.ID)

	if activeCurrentTask != nil {
		coins := s.getCurrentTaskCoins(activeCurrentTask.ID)
		activeCurrentTask.Coins = &coins
		s.CurrentTask.UpdateCurrentTaskById(activeCurrentTask)
	}

	currentTasks, err := s.CurrentTask.GetCurrentTasksByUserId(user.ID)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return currentTasks, nil
}
