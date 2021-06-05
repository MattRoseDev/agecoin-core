package database

import (
	"fmt"
	"time"

	"github.com/favecode/agecoin-core/graph/model"
	"github.com/go-pg/pg"
)

type Task struct {
	DB *pg.DB
}

func (t *Task) GetTaskByField(field, value string) (*model.Task, error) {
	var task model.Task
	err := t.DB.Model(&task).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	return &task, err
}

func (t *Task) GetTaskByID(id string) (*model.Task, error) {
	return t.GetTaskByField("id", id)
}

func (t *Task) CreateTask(task *model.Task) (*model.Task, error) {
	_, err := t.DB.Model(task).Returning("*").Insert()
	return task, err
}

func (t *Task) GetTasksByUserId(userId string) ([]*model.Task, error) {
	var tasks []*model.Task
	err := t.DB.Model(&tasks).Where("user_id = ?" ,userId).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*").Select()
	return tasks, err
}

func (t *Task) UpdateTaskById(task *model.Task) (*model.Task, error) {
	_, err := t.DB.Model(task).Where("id = ?" ,task.ID).Where("deleted_at is ?", nil).Returning("*").Update()
	return task, err
}

func (t *Task) DeleteTaskById(taskId string) (*model.Task, error) {
	DeletedAt := time.Now()
	var task = &model.Task{
		ID: taskId,
		DeletedAt: &DeletedAt, 
	}
	_, err := t.DB.Model(task).Set("deleted_at = ?deleted_at").Where("id = ?id").Where("deleted_at is ?", nil).Returning("*").Update()
	return task, err
}