package model

import (
	"time"
)

type CurrentTask struct {
	tableName    struct{}   `sql:"current_task"`
	ID           string     `json:"id"`
	UserID       string     `json:"userId"`
	TaskID       string     `json:"taskId"`
	Description  *string    `json:"description"`
	DefaultCoins int        `json:"defaultCoins"`
	Coins        *int       `json:"coins"`
	Status       int        `json:"status"`
	Active       *bool      `json:"active"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"-" pg:",soft_delete"`
}
