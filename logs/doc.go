// Package logs provides high-level operations for managing log records:
// creating new logs, editing existing ones, deleting logs, and listing/opening
// logs for interactive use.
//
// # Database-Agnostic Design
//
// This package does not implement persistence and does not assume any
// particular storage technology. It delegates data access and mutation to a
// database layer that exposes four operations: Get, Create, Patch, and Delete.
// Whether that layer is in-memory, file-backed, SQL, or NoSQL is irrelevant to
// this package.
//
// # Presentation Index
//
// The package maintains an ordered index of log IDs to present a stable,
// 1-based menu/order to callers. Functions validate indices, perform simple
// input checks, and then call the database layer to read or mutate log data.
//
// # Concurrency
//
// The package is intended for single-goroutine CLI-style use and does not
// provide synchronization. If used concurrently, callers must synchronize
// access to package-level state.
//
// # Errors and Prompts
//
// Functions intended for interactive flows surface invalid indices and missing
// records via user-facing prompts. Functions that load data from external
// sources (if present) return errors from the underlying database loader.
//
// # Usage
//
// The package is typically driven by a CLI loop:
//
//	// Create a new log (interactive).
//	logs.Create_log()
//
//	// Edit the first log (index 0).
//	logs.Edit_log(0)
//
//	// Delete the first log.
//	logs.Delete_log(0)
//
//	// List or open logs for viewing/editing.
//	logs.Print_logs()
//	logs.Open_log(0)
//
// Package logs.
package logs
