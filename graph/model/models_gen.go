// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type AddTaskInput struct {
	Title        string  `json:"title"`
	Description  *string `json:"description"`
	DefaultCoins int     `json:"defaultCoins"`
}

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *User      `json:"user"`
}

type AuthToken struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}

type DailyCoins struct {
	RemainingCoins int   `json:"remainingCoins"`
	SavedCoins     int   `json:"savedCoins"`
	WastedCoins    int   `json:"wastedCoins"`
	ActiveTask     *Task `json:"activeTask"`
}

type EditTaskInput struct {
	Title        *string `json:"title"`
	Description  *string `json:"description"`
	DefaultCoins *int    `json:"defaultCoins"`
	Coins        *int    `json:"coins"`
}

type FinishTaskInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Coins       *int    `json:"coins"`
}

type GetTasksFilter struct {
	Status *int `json:"status"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Fullname *string `json:"fullname"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

type Test struct {
	Content *string `json:"content"`
}

type TestInput struct {
	Content *string `json:"content"`
}
