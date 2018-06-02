package pkg

import (
	"time"
)

var (
	// SecondsOfDay is number of seconds of one day
	SecondsOfDay = secondsOfDuration(time.Hour * 24)
)

func secondsOfDuration(duration time.Duration) int64 {
	println("calculating seconds of duration ")
	return int64(duration.Seconds())
}
