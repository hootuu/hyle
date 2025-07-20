package htimes

import "time"

func NxtYearDay(current time.Time) (startTime, EndTime time.Time) {
	startTime = time.Date(
		current.Year(),
		current.Month(),
		current.Day(),
		0, 0, 0, 0,
		current.Location(),
	)

	nextYearSameDay := startTime.AddDate(1, 0, 0)

	EndTime = time.Date(
		nextYearSameDay.Year(),
		current.Month(),
		current.Day(),
		0, 0, 0, 0,
		current.Location(),
	).Add(-time.Second)

	return startTime, EndTime
}
