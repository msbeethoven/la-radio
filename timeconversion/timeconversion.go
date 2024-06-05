package timeconversion

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseDurationFromHHMMSS(durationString string) (time.Duration, error) {
	parts := strings.Split(durationString, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid duration format: %s", durationString)
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, err
	}

	return time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second, nil
}
