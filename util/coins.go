package util

import (
	"time"
)

func CalculateCurrentTaskCoins(t time.Time) int {
	now := time.Now().Unix()
	result := now - t.Unix()

	return int(result / 60)
}
