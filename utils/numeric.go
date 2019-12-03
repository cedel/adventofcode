package utils

import "strconv"

// ToInt converts a given string to an int
func ToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
