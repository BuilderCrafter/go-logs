package helpers

import (
	"fmt"
	"os"
	"strings"
)

func ClearScreen() {
	// Clears the screen and moves the cursor to home
	ansiClear := "\033[2J\033[H"
	if _, err := fmt.Fprint(os.Stdout, ansiClear); err == nil {
		return
	}
}

func WelcomeMessage() {
	fmt.Println("Welcome to logs storage application!")
	fmt.Println("This application allows the storage of text logs")
	fmt.Println("You can read, write, edit and delete logs")
	fmt.Println("Enter the number of log you would like to read/edit/delete")
	fmt.Println("Enter '+' if you would like to create new log")
	fmt.Println("---------------------LOGS------------------------")
}

func WrapText(s string, limit int) string {
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
