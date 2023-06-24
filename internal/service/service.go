package service

import (
	"time"
)

var timeUntil = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

// DaysLeft считает кол-во дней, которое осталось до 1 января 2025
func DaysLeft() int {
	timeLeft := time.Until(timeUntil)
	daysLeft := int(timeLeft.Hours() / 24)
	return daysLeft
}
