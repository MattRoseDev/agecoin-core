package service

import (
	"context"
	"errors"

	"github.com/favecode/agecoin-core/graph/model"
	"github.com/favecode/agecoin-core/middleware"
	"github.com/favecode/agecoin-core/util"
)

func (s *Service) saveTaskHistory(task *model.Task, historyType string) {
	taskHistory := &model.TaskHistory{
		UserID: task.UserID,
		TaskID: task.ID,
		Type:   historyType,
	}

	s.TaskHistory.CreateTaskHistory(taskHistory)
}

func (s *Service) getTaskCoins(taskID string) int {
	taskHistory, _ := s.TaskHistory.GetTaskHistoryByTaskId(taskID)

	if len(taskHistory.ID) < 1 {
		return 0
	} else {
		return util.CalculateTaskCoins(taskHistory.CreatedAt)
	}
}

func (s *Service) AddTask(ctx context.Context, input model.AddTaskInput) (*model.Task, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	task := &model.Task{
		UserID:       user.ID,
		Title:        input.Title,
		Description:  input.Description,
		DefaultCoins: input.DefaultCoins,
	}

	s.Task.CreateTask(task)

	return task, nil
}

func (s *Service) EditTask(ctx context.Context, taskID string, input model.EditTaskInput) (*model.Task, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	task, err := s.Task.GetTaskByID(taskID)

	if err != nil {
		return nil, errors.New("task not found")
	}

	if task.UserID != user.ID {
		return nil, errors.New("authorization failed")
	}

	didUpdate := false

	if input.Title != nil {
		task.Title = *input.Title
		didUpdate = true
	}

	if input.Description != nil {
		task.Description = input.Description
		didUpdate = true
	}

	if input.DefaultCoins != nil {
		task.DefaultCoins = *input.DefaultCoins
		didUpdate = true
	}

	if input.Coins != nil {
		task.Coins = *input.Coins
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update done")
	}

	newTask, err := s.Task.UpdateTaskById(task)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return newTask, nil
}

func (s *Service) StartTask(ctx context.Context, taskID string) (*model.Task, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	task, _ := s.Task.GetTaskByID(taskID)

	if len(task.ID) < 1 {
		return nil, errors.New("task not found")
	}

	if task.Active == bool(true) {
		return nil, errors.New("task is active")
	}

	s.Task.DeactiveAllTaskByUserId(user.ID)

	task.Active = true
	task.Status = 1

	newTask, err := s.Task.UpdateTaskById(task)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	s.saveTaskHistory(task, "START")

	return newTask, nil
}

func (s *Service) PauseTask(ctx context.Context, taskID string) (*model.Task, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	task, _ := s.Task.GetTaskByID(taskID)

	if len(task.ID) < 1 {
		return nil, errors.New("task not found")
	}

	if task.Active == bool(false) {
		return nil, errors.New("task is not active")
	}

	taskHistory, _ := s.TaskHistory.GetTaskHistoryByTaskId(taskID)

	coins := util.CalculateTaskCoins(taskHistory.CreatedAt)

	task.Active = bool(false)
	task.Status = 1
	task.Coins = coins

	newTask, err := s.Task.UpdateTaskById(task)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	s.Task.DeactiveAllTaskByUserId(user.ID)

	s.saveTaskHistory(task, "PAUSE")

	return newTask, nil
}

func (s *Service) FinishTask(ctx context.Context, taskID string, input *model.FinishTaskInput) (*model.Task, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	task, _ := s.Task.GetTaskByID(taskID)

	if len(task.ID) < 1 {
		return nil, errors.New("task not found")
	}

	if task.Status == 2 && input == nil {
		return task, nil
	}

	if task.Active == bool(true) {
		if input.Coins == nil {
			coins := s.getTaskCoins(task.ID) + task.Coins
			task.Coins = coins
		}
		task.Active = bool(false)
	}

	task.Status = 2

	if input.Title != nil {
		task.Title = *input.Title
	}

	if input.Description != nil {
		task.Description = input.Description
	}

	if input.Coins != nil {
		task.Coins = *input.Coins
	}

	newTask, err := s.Task.UpdateTaskById(task)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	s.Task.DeactiveAllTaskByUserId(user.ID)

	s.saveTaskHistory(task, "FINISH")

	return newTask, nil
}

func (s *Service) DeleteTask(ctx context.Context, taskID string) (*model.Task, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	task, _ := s.Task.GetTaskByID(taskID)

	if len(task.ID) < 1 {
		return nil, errors.New("task not found")
	}

	if task.UserID != user.ID {
		return nil, errors.New("authorization failed")
	}

	deletedTask, err := s.Task.DeleteTaskById(taskID)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	s.saveTaskHistory(task, "CANCEL")

	return deletedTask, nil
}

func (s *Service) GetTasks(ctx context.Context, filter *model.GetTasksFilter) ([]*model.Task, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	activeTask, _ := s.Task.GetActiveTaskByUserId(user.ID)

	if activeTask != nil {
		coins := s.getTaskCoins(activeTask.ID)
		activeTask.Coins = coins
		s.Task.UpdateTaskById(activeTask)
	}

	tasks, err := s.Task.GetTasksByUserId(user.ID, filter)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return tasks, nil
}

func (s *Service) GetTask(ctx context.Context, taskID string) (*model.Task, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	task, err := s.Task.GetTaskByUserIdAndID(user.ID, taskID)

	if len(task.ID) == 0 {
		return nil, errors.New("task not found")
	}

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return task, nil
}
