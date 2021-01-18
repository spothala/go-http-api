package utils

import (
	"errors"
	"strconv"
)

// Duration - Time Duration to retrieve logs
// age can be 30m, 1h, 7d
func Duration(age string) (int, error) {
	duration, err := strconv.Atoi(age[:len(age)-1])
	if err != nil {
		return 0, err
	}
	totalSeconds := 0
	switch age[len(age)-1:] {
	case "m":
		totalSeconds = duration * 60
	case "h":
		totalSeconds = duration * 60 * 60
	case "d":
		totalSeconds = duration * 60 * 60 * 24
	default:
		err = errors.New("Duration should be in requested format. e.g. 1m,1h,1d")
	}
	return totalSeconds, err
}
