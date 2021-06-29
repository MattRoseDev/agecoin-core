package model

import (
	"time"
)

type Task struct {
	tableName    struct{}   `sql:"task"`
	ID           string     `json:"id"`
	UserID       string     `json:"userId"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	DefaultCoins int        `json:"defaultCoins"`
	Coins        *int       `json:"coins"`
	Status       int        `json:"status"`
	Active       bool       `json:"active"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"-" pg:",soft_delete"`
}
