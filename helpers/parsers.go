package helpers

import (
	"strconv"
	"strings"
)

func LogParseInput(s string) (str string, valid bool) {
	// Valid inputs for Log
	valid_inputs := map[string]bool{
		"x": true,
		"e": true,
		"r": true,
	}
	if valid_inputs[strings.ToLower(s)] {
		return s, true
	}
	return "", false

}

func MenuParseInput(s string) (num int, str string) {
	s = strings.TrimSpace(s)

	if n, err := strconv.Atoi(s); err == nil && n > 0 {
		return n, ""
	}
	return 0, s
}
