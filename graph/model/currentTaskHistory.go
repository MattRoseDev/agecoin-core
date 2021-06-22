package model

import (
	"time"
)

type CurrentTaskHistory struct {
	tableName struct{}   `sql:"current_task_history"`
	ID        string     `json:"id"`
	UserID    string     `json:"userId"`
	TaskID    string     `json:"taskId"`
	Type      *string    `json:"type"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" pg:",soft_delete"`
}