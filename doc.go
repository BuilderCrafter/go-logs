// Command go-logs starts an interactive, terminal-based log editor.
//
// The program presents a simple menu to create, view, edit, and delete logs.
// Storage and persistence are delegated to a separate database layer; this
// command does not depend on any particular database technology.
//
// Usage:
//
//	go-logs [-load file.json]
//
// Controls:
//
//	<number>   Open a log by its menu number (1-based).
//	+          Create a new log.
//	q          Quit the application.
//
// In the "open log" view:
//
//	e          Edit the current log.
//	x          Delete the current log.
//	r          Return to the list.
//
// Flags:
//
//	-load file.json   Load logs from a JSON file on startup (optional).
//
// The program reads from standard input and writes to standard output. It is
// intended for single-goroutine, interactive use.
//
// Command go-logs.
package main
