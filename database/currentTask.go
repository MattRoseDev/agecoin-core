package database

import (
	"github.com/favecode/agecoin-core/graph/model"
	"github.com/go-pg/pg"
)

type CurrentTask struct {
	DB *pg.DB
}

func (c *CurrentTask) CreateCurrentTask(currentTask *model.CurrentTask) (*model.CurrentTask, error) {
	_, err := c.DB.Model(currentTask).Returning("*").Insert()
	return currentTask, err
}
