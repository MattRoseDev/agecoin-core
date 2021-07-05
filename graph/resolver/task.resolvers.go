package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/favecode/agecoin-core/graph/model"
)

func (r *mutationResolver) AddTask(ctx context.Context, input model.AddTaskInput) (*model.Task, error) {
	return r.Service.AddTask(ctx, input)
}

func (r *mutationResolver) EditTask(ctx context.Context, taskID string, input *model.EditTaskInput) (*model.Task, error) {
	return r.Service.EditTask(ctx, taskID, input)
}

func (r *mutationResolver) DeleteTask(ctx context.Context, taskID string) (*model.Task, error) {
	return r.Service.DeleteTask(ctx, taskID)
}

func (r *mutationResolver) StartTask(ctx context.Context, taskID string) (*model.Task, error) {
	return r.Service.StartTask(ctx, taskID)
}

func (r *mutationResolver) PauseTask(ctx context.Context, taskID string) (*model.Task, error) {
	return r.Service.PauseTask(ctx, taskID)
}

func (r *mutationResolver) FinishTask(ctx context.Context, taskID string, input *model.FinishTaskInput) (*model.Task, error) {
	return r.Service.FinishTask(ctx, taskID, input)
}

func (r *queryResolver) GetTasks(ctx context.Context, filter *model.GetTasksFilter) ([]*model.Task, error) {
	return r.Service.GetTasks(ctx, filter)
}

func (r *queryResolver) GetTask(ctx context.Context, taskID string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}
