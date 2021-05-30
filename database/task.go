package database

import (
	"fmt"

	"github.com/favecode/agecoin-core/graph/model"
	"github.com/go-pg/pg"
)

type Task struct {
	DB *pg.DB
}

func (t *Task) GetTaskByField(field, value string) (*model.Task, error) {
	var user model.Task
	err := t.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	return &user, err
}

func (t *Task) GetTaskByID(id string) (*model.Task, error) {
	return t.GetTaskByField("id", id)
}

func (t *Task) CreateTask(user *model.Task) (*model.Task, error) {
	_, err := t.DB.Model(user).Returning("*").Insert()
	return user, err
}
