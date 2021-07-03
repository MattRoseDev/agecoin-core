package database

import (
	"github.com/favecode/agecoin-core/graph/model"
	"github.com/go-pg/pg"
)

type TaskHistory struct {
	DB *pg.DB
}

func (c *TaskHistory) CreateTaskHistory(taskHistory *model.TaskHistory) (*model.TaskHistory, error) {
	_, err := c.DB.Model(taskHistory).Returning("*").Insert()
	return taskHistory, err
}

func (c *TaskHistory) GetTaskHistoryByTaskIdAndType(taskId string, types []string) ([]*model.TaskHistory, error) {
	var taskHistory []*model.TaskHistory
	err := c.DB.Model(&taskHistory).Where("task_id = ?", taskId).Where("type in (?)", pg.In(types)).Where("deleted_at is ?", nil).Order("created_at ASC").Select()
	return taskHistory, err
}
