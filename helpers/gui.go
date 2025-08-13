// Package helpers provides small CLI utilities for screen clearing,
// welcome/prompt output, line input, and word-wrapping for text display.
//
// These helpers are designed for simple, single-goroutine, terminal-based
// programs. They write to os.Stdout and read from os.Stdin by default.
package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// stdin is the reader used by ReadLine and WaitForEnter. It defaults to os.Stdin
var stdin = bufio.NewReader(os.Stdin)

// ClearScreen clears the terminal screen and moves the cursor to the home
// position using ANSI escape sequences. On terminals that do not support
// ANSI (or if writing fails), the function simply returns without error.
//
// Note: On modern Windows terminals and WSL, ANSI is generally supported.
// If broader compatibility is needed, consider a platform-specific fallback.
func ClearScreen() {
	ansiClear := "\033[2J\033[H"
	if _, err := fmt.Fprint(os.Stdout, ansiClear); err == nil {
		return
	}
}

// WelcomeMessage prints a short description of the CLI and its basic commands.
func WelcomeMessage() {
	fmt.Println("Welcome to logs storage application!")
	fmt.Println("This application allows the storage of text logs")
	fmt.Println("You can read, write, edit and delete logs")
	fmt.Println("Enter the number of log you would like to read/edit/delete")
	fmt.Println("Enter '+' if you would like to create new log")
	fmt.Println("---------------------LOGS------------------------")
}

// WrapText wraps s at word boundaries so that no line exceeds limit runes,
// except when a single word is longer than limit. In that case, the long
// word is placed on its own line without hyphenation. Whitespace is
// normalized: multiple spaces/tabs/newlines collapse to single spaces.
//
// If limit <= 0, WrapText returns s unchanged.
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

// ReadLine reads a single line from stdin, trimming the trailing "\r\n" or
// "\n". On I/O error it returns what was read up to that point (if any),
// with line endings trimmed.
func ReadLine() string {
	s, _ := stdin.ReadString('\n')
	return strings.TrimRight(s, "\r\n")
}

// WaitForEnter prints prompt (or a default message if prompt is empty) and
// then waits for the user to press Enter. Any extra characters typed on
// the same line are discarded.
func WaitForEnter(prompt string) {
	if prompt == "" {
		fmt.Print("-Press Enter to continue-")
	} else {
		fmt.Println(prompt)
	}
	stdin.ReadString('\n')
}
