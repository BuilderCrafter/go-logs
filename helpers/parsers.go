package helpers

import (
	"strconv"
	"strings"
)

// logActions is the set of valid single-letter actions in the "open log" view.
// Kept as a package-level var to avoid re-allocating a map on each parse call.
var logActions = map[string]struct{}{
	"x": {}, // delete
	"e": {}, // edit
	"r": {}, // return
}

// ParseLogInput validates an action entered in the "open log" view.
//
// It trims surrounding whitespace, lowercases the input, and returns the
// normalized action plus a boolean indicating validity.
// Accepted actions: "x" (delete), "e" (edit), "r" (return).
//
// Examples:
//
//	ParseLogInput("E")   => "e", true
//	ParseLogInput("  x") => "x", true
//	ParseLogInput("q")   => "",  false
func LogParseInput(s string) (str string, valid bool) {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	_, ok := logActions[s]
	if ok {
		return s, true
	}
	return "", false

}

// ParseMenuInput distinguishes between a positive menu number and a non-numeric command.
//
// It trims whitespace. If s is a positive integer (e.g., "1", "42"), it returns
// num>0 and an empty command (""). Otherwise it returns num==0 and the
// lowercased command string (e.g., "+", "q", "help").
//
// This keeps the caller logic simple:
//
//	if num > 0 { /* user chose an item */ } else { /* handle command */ }
//
// Examples:
//
//	ParseMenuInput("7")    => 7, ""
//	ParseMenuInput("+")    => 0, "+"
//	ParseMenuInput(" Help ") => 0, "help"
func ParseMenuInput(s string) (num int, cmd string) {
	s = strings.TrimSpace(s)

	if n, err := strconv.Atoi(s); err == nil && n > 0 {
		return n, ""
	}
	return 0, strings.ToLower(s)
}
