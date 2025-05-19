package main

import (
	"fmt"
)

func ConvertToDigitalFormat(hour, minute, second int) string {
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func ExtractTimeUnits(seconds int) (int, int, int) {
	hours := seconds / 3600
	seconds %= 3600
	minutes := seconds / 60
	seconds %= 60

	return hours, minutes, seconds
}
