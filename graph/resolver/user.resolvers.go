package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/favecode/agecoin-core/graph/model"
)

func (r *queryResolver) GetUserInfo(ctx context.Context) (*model.User, error) {
	return r.Service.GetUserInfo(ctx)
}

func (r *queryResolver) GetDailyCoins(ctx context.Context) (*model.DailyCoins, error) {
	return r.Service.GetDailyCoins(ctx)
}
