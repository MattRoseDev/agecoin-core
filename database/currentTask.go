package database

import (
	"fmt"
	"time"

	"github.com/favecode/agecoin-core/graph/model"
	"github.com/go-pg/pg"
)

type CurrentTask struct {
	DB *pg.DB
}

func (c *CurrentTask) GetCurrentTaskByField(field, value string) (*model.CurrentTask, error) {
	var currentTask model.CurrentTask
	err := c.DB.Model(&currentTask).Where(fmt.Sprintf("%v = ?", field), value).Where("deleted_at is ?", nil).First()
	return &currentTask, err
}

func (c *CurrentTask) GetCurrentTaskByID(id string) (*model.CurrentTask, error) {
	return c.GetCurrentTaskByField("id", id)
}

func (c *CurrentTask) GetCurrentTasksByUserId(userId string) ([]*model.CurrentTask, error) {
	var currentTasks []*model.CurrentTask
	err := c.DB.Model(&currentTasks).Where("user_id = ?", userId).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*").Select()
	return currentTasks, err
}

func (c *CurrentTask) GetActiveCurrentTaskByUserId(userId string) (*model.CurrentTask, error) {
	var currentTask model.CurrentTask
	err := c.DB.Model(&currentTask).Where("user_id = ?", userId).Where("active = ?", true).Where("deleted_at is ?", nil).Order("created_at DESC").Returning("*").Select()
	return &currentTask, err
}

func (c *CurrentTask) CreateCurrentTask(currentTask *model.CurrentTask) (*model.CurrentTask, error) {
	_, err := c.DB.Model(currentTask).Returning("*").Insert()
	return currentTask, err
}

func (c *CurrentTask) DeactiveAllCurrentTaskByUserId(userId string) ([]*model.CurrentTask, error) {
	var currentTasks []*model.CurrentTask
	_, err := c.DB.Model(&currentTasks).Set("active = ?", false).Where("user_id = ?", userId).Where("deleted_at is ?", nil).Returning("*").Update()
	return currentTasks, err
}

func (c *CurrentTask) UpdateCurrentTaskById(currentTask *model.CurrentTask) (*model.CurrentTask, error) {
	_, err := c.DB.Model(currentTask).Where("id = ?", currentTask.ID).Where("deleted_at is ?", nil).Returning("*").Update()
	return currentTask, err
}

func (c *CurrentTask) DeleteCurrentTaskById(currentTaskId string) (*model.CurrentTask, error) {
	DeletedAt := time.Now()
	var currentTask = &model.CurrentTask{
		ID:        currentTaskId,
		DeletedAt: &DeletedAt,
	}
	_, err := c.DB.Model(currentTask).Set("deleted_at = ?deleted_at").Where("id = ?id").Where("deleted_at is ?", nil).Returning("*").Update()
	return currentTask, err
}
