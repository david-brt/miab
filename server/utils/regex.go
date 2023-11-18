package utils

import (
	"regexp"
)

// helper function that iterates ofer multiple regex conditions and returns whether all of the match the given string
func checkRegexes(regexes []*regexp.Regexp, pattern string) bool {
	for _, regex := range regexes {
		if !regex.MatchString(pattern) {
			return false
		}
	}
	return true
}

// returns whether a given password matches the regular expression ^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$ %^&*-]).{8,}$
func IsValidPassword(password string) bool {
	// regex must be split as regexp does not support lookaheads
	upperCase := regexp.MustCompile(`[A-Z]`)
	lowerCase := regexp.MustCompile(`[a-z]`)
	digit := regexp.MustCompile(`[0-9]`)
	specialChar := regexp.MustCompile(`[#?!@$ %^&*]`)

	regexes := []*regexp.Regexp{
		upperCase,
		lowerCase,
		digit,
		specialChar,
	}

	if len(password) < 8 || len(password) > 50 {
		return false
	}

	return checkRegexes(regexes, password)
}

func IsValidUsername(username string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9_]{1,20}$`)
	return regex.MatchString(username)
}
