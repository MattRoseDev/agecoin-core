package database

import (
	"github.com/favecode/agecoin-core/graph/model"
	"github.com/go-pg/pg"
)

type CurrentTaskHistory struct {
	DB *pg.DB
}

func (c *CurrentTaskHistory) CreateCurrentTaskHistory(currentTaskHistory *model.CurrentTaskHistory) (*model.CurrentTaskHistory, error) {
	_, err := c.DB.Model(currentTaskHistory).Returning("*").Insert()
	return currentTaskHistory, err
}

func (c *CurrentTaskHistory) GetCurrentTaskHistoryByCurrentTaskId(currentTaskId string) (*model.CurrentTaskHistory, error) {
	var currentTaskHistory model.CurrentTaskHistory
	err := c.DB.Model(&currentTaskHistory).Where("current_task_id = ?", currentTaskId).Where("deleted_at is ?", nil).Order("created_at DESC").First()
	return &currentTaskHistory, err
}
