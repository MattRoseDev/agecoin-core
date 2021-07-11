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

func (c *Task) GetTaskByField(field, value string) (*model.Task, error) {
	var task model.Task
	err := c.DB.Model(&task).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	return &task, err
}

func (c *Task) GetTaskByID(id string) (*model.Task, error) {
	return c.GetTaskByField("id", id)
}

func (c *Task) GetTasksByUserId(userId string, filter *model.GetTasksFilter) ([]*model.Task, error) {
	var tasks []*model.Task
	query := c.DB.Model(&tasks).Where("user_id = ?", userId).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*")

	if filter != nil {
		if filter.Status != nil {
			query.Where("status = ?", filter.Status)
		} else {
			query.Where("status < ?", 3)
		}

		if filter.Daily != nil && *filter.Daily == bool(true) {
			if filter.TimezoneOffset != nil {
				now := time.Now().UTC()
				start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Add(time.Duration(-*filter.TimezoneOffset) * time.Minute)
				end := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()).Add(time.Duration(-*filter.TimezoneOffset) * time.Minute)

				query.Where("created_at >= ?", start).Where("created_at <= ?", end)
			}
		}
	}

	err := query.Select()

	return tasks, err
}

func (t *Task) GetTaskByUserIdAndID(userId string, id string) (*model.Task, error) {
	var task model.Task
	err := t.DB.Model(&task).Where("id = ?", id).Where("user_id = ?", userId).Where("deleted_at is ?", nil).First()
	return &task, err
}

func (c *Task) GetActiveTaskByUserId(userId string) (*model.Task, error) {
	var task model.Task
	err := c.DB.Model(&task).Where("user_id = ?", userId).Where("active = ?", true).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*").Select()
	if len(task.ID) < 1 {
		return nil, nil
	}
	return &task, err
}

func (c *Task) CreateTask(task *model.Task) (*model.Task, error) {
	_, err := c.DB.Model(task).Returning("*").Insert()
	return task, err
}

func (c *Task) DeactiveTaskByUserIdAndTaskId(userId string, taskId string) ([]*model.Task, error) {
	var tasks []*model.Task
	_, err := c.DB.Model(&tasks).Set("active = ?", false).Where("user_id = ?", userId).Where("id = ?", taskId).Where("deleted_at is ?", nil).Returning("*").Update()
	return tasks, err
}

func (c *Task) UpdateTaskById(task *model.Task) (*model.Task, error) {
	_, err := c.DB.Model(task).Where("id = ?", task.ID).Where("deleted_at is ?", nil).Returning("*").Update()
	return task, err
}

func (c *Task) DeleteTaskById(taskId string) (*model.Task, error) {
	DeletedAt := time.Now()
	var task = &model.Task{
		ID:        taskId,
		DeletedAt: &DeletedAt,
	}
	_, err := c.DB.Model(task).Set("deleted_at = ?deleted_at").Where("id = ?id").Where("deleted_at is ?", nil).Returning("*").Update()
	return task, err
}
