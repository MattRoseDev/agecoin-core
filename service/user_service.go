package service

import (
	"context"
	"errors"
	"time"

	"github.com/favecode/agecoin-core/graph/model"
	"github.com/favecode/agecoin-core/middleware"
)

func (s *Service) GetUserInfo(ctx context.Context) (*model.User, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	return user, nil
}

func (s *Service) GetDailyCoins(ctx context.Context, input model.InputGetDailyCoins) (*model.DailyCoins, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	daily := bool(true)

	tasks, _ := s.GetTasks(ctx, &model.GetTasksFilter{
		Daily:          &daily,
		TimezoneOffset: &input.TimezoneOffset,
	})

	savedCoins := 0

	for _, t := range tasks {
		savedCoins += t.Coins
	}

	spentCoins := (time.Now().Hour()*60 + time.Now().Minute()) * 60

	remainingCoins := (60 * 60 * 24) - spentCoins

	activeTask, _ := s.Task.GetActiveTaskByUserId(user.ID)

	return &model.DailyCoins{
		SavedCoins:     savedCoins,
		RemainingCoins: remainingCoins,
		WastedCoins:    (spentCoins - savedCoins),
		ActiveTask:     activeTask,
	}, nil
}
