package logs

import "strings"

func wrap(s string, limit int) string {
	if limit <= 0 {
		return s
	}
	words := strings.Fields(s) // collapses all whitespace to single spaces
	var b strings.Builder
	lineLen := 0

	for i, w := range words {
		if lineLen+len(w) > limit { // word wouldn’t fit → new line
			b.WriteByte('\n')
			lineLen = 0
		} else if i != 0 { // not first word on a line → add space
			b.WriteByte(' ')
			lineLen++
		}
		b.WriteString(w)
		lineLen += len(w)
	}
	return b.String()
}
