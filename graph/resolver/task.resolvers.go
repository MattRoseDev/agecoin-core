package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/favecode/agecoin-core/graph/model"
)

func (r *mutationResolver) AddTask(ctx context.Context, input model.AddTaskInput) (*model.Task, error) {
	return r.Service.AddTask(ctx, input)
}

func (r *mutationResolver) EditTask(ctx context.Context, taskID string, input model.EditTaskInput) (*model.Task, error) {
	return r.Service.EditTask(ctx, taskID, input)
}

func (r *mutationResolver) DeleteTask(ctx context.Context, taskID string) (*model.Task, error) {
	return r.Service.DeleteTask(ctx, taskID)
}

func (r *queryResolver) GetTasks(ctx context.Context) ([]*model.Task, error) {
	return r.Service.GetTasks(ctx)
}

func (r *queryResolver) GetTask(ctx context.Context, taskID string) (*model.Task, error) {
	return r.Service.GetTask(ctx, taskID)
}
