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
	DefaultCoins int        `json:"defaultCoins" sql:",notnull"`
	Coins        int        `json:"coins" sql:",notnull"`
	Status       int        `json:"status" sql:",notnull"`
	Active       bool       `json:"active" sql:",use_zero,notnull"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"-" sql:",soft_delete"`
}
