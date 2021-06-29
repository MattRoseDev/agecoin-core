package util

import (
	"time"
)

func CalculateTaskCoins(t time.Time) int {
	now := time.Now().Unix()
	result := now - t.Unix()

	return int(result / 60)
}
