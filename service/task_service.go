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
		task.Title= *input.Title
		didUpdate = true
	}

	if input.Description != nil {
		task.Description = input.Description
		didUpdate = true
	}

	if input.Description != nil {
		task.DefaultCoins = *input.DefaultCoins
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


func (s *Service) DeleteTask(ctx context.Context, taskID string) (*model.Task, error) {
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

	deletedTask, err := s.Task.DeleteTaskById(task.ID)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	
	return deletedTask, nil
}