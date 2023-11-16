package utils

import "time"

func CurrentYear() int {
	return time.Now().Year()
}

func YearsAgo(years int) int {
	return time.Now().Year() - years

}
