package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/favecode/agecoin-core/graph/model"
)

func (r *mutationResolver) AddCurrentTask(ctx context.Context, input model.AddCurrentTaskInput) (*model.CurrentTask, error) {
	return r.Service.AddCurrentTask(ctx, input)
}

func (r *mutationResolver) EditCurrentTask(ctx context.Context, currentTaskID string, input model.EditCurrentTaskInput) (*model.CurrentTask, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	return r.Service.DeleteCurrentTask(ctx, currentTaskID)
}

func (r *mutationResolver) StartCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	return r.Service.StartCurrentTask(ctx, currentTaskID)
}

func (r *mutationResolver) PauseCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	return r.Service.PauseCurrentTask(ctx, currentTaskID)
}

func (r *mutationResolver) FinishCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	return r.Service.FinishCurrentTask(ctx, currentTaskID)
}

func (r *queryResolver) GetCurrentTasks(ctx context.Context) ([]*model.CurrentTask, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	panic(fmt.Errorf("not implemented"))
}
