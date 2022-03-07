package util

import (
	"regexp"
	"strings"
)

const minLenght = 8
const maxLength = 32

//StrongPassword check password polciy
func StrongPassword(password string) bool {
	var (
		count = 0
	)

	if len(password) >= minLenght && len(password) <= maxLength {
		count++
	}

	match, _ := regexp.MatchString(".*[a-z].*", password)
	if match == true {
		count++
	}

	match, _ = regexp.MatchString(".*[A-Z].*", password)
	if match == true {
		count++
	}

	match, _ = regexp.MatchString(`.*\W.*`, password)
	if match == true {
		count++
	}

	match, _ = regexp.MatchString(".*[0-9].*", password)
	if match == true {
		count++
	}
	match = strings.Contains(password, " ")
	if match == false {
		count++
	}
	return count > 5
}
