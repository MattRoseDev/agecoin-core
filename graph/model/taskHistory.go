package model

import (
	"time"
)

type TaskHistory struct {
	tableName struct{}   `sql:"task_history"`
	ID        string     `json:"id"`
	UserID    string     `json:"userId"`
	TaskID    string     `json:"TaskId"`
	Type      string     `json:"type"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" pg:",soft_delete"`
}
