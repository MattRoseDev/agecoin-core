package database

import (
	"github.com/favecode/agecoin-core/graph/model"
	"github.com/go-pg/pg"
)

type CurrentTaskHistory struct {
	DB *pg.DB
}

func (c *CurrentTaskHistory) CreateCurrentTaskHistory(currentTask *model.CurrentTaskHistory) (*model.CurrentTaskHistory, error) {
	_, err := c.DB.Model(currentTask).Returning("*").Insert()
	return currentTask, err
}
