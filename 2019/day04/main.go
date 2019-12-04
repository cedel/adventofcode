package main

import "fmt"

const (
	passwordLength = 6
	min            = 172930
	max            = 683082
)

// ValidPassword determines if the given password sattisfies
// all of the following criteria:
// 1. It must be a 6 digit number
// 2. Each digit must be greater than or equal to the previous one
// 3. There must be at least one pair of consecutive digits
func ValidPassword(input int) bool {
	s := []byte(fmt.Sprintf("%d", input))
	if len(s) != passwordLength {
		return false
	}

	doubleDigit := false
	for i := 1; i < passwordLength; i++ {
		if s[i] < s[i-1] {
			return false
		}

		if !doubleDigit && s[i] == s[i-1] {
			doubleDigit = true
		}
	}

	return doubleDigit
}

// ValidStrictPassword determines if the given password sattisfies
// all of the following criteria:
// 1. It must be a 6 digit number
// 2. Each digit must be greater than or equal to the previous one
// 3. There must be at least one pair of consecutive digits that are
//    not part of a larger sequence, i.e. exactly two consecutive digits
func ValidStrictPassword(input int) bool {
	s := []byte(fmt.Sprintf("%d", input))
	if len(s) != passwordLength {
		return false
	}

	doubleDigit := false
	for i := 1; i < passwordLength; i++ {
		if s[i] < s[i-1] {
			return false
		}

		if !doubleDigit && (i == 1 || s[i-2] != s[i-1]) && s[i-1] == s[i] && (i == passwordLength-1 || s[i] != s[i+1]) {
			doubleDigit = true
		}
	}

	return doubleDigit
}

func main() {
	validPasswords := 0
	validStrictPasswords := 0
	for i := min; i < max; i++ {
		if ValidPassword(i) {
			validPasswords++
		}
		if ValidStrictPassword(i) {
			validStrictPasswords++
		}
	}
	fmt.Println(validPasswords, validStrictPasswords)
}
