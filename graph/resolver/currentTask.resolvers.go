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
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StartCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PauseCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) FinishCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCurrentTasks(ctx context.Context) ([]*model.CurrentTask, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	panic(fmt.Errorf("not implemented"))
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) StopCurrentTask(ctx context.Context, currentTaskID string) (*model.CurrentTask, error) {
	panic(fmt.Errorf("not implemented"))
}