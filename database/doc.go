// Package database provides an in-memory store for log records and the Log
// model type. It exposes simple CRUD-style helpers—Get, Create, Patch, and
// Delete—and a loader that imports logs from a JSON file.
//
// # Overview
//
// Data is held in a package-level map keyed by uint64 IDs. New IDs are issued
// monotonically via an atomic counter. The package itself does not persist data;
// callers that need durability should add their own persistence layer.
//
// # JSON Import
//
// LoadFromJSON reads a file containing an array of Log objects and inserts them
// into the store, assigning fresh IDs for each record. Any ID in the input is
// ignored (the Log.ID field is not serialized/deserialized).
//
// # Concurrency
//
// This package is intended for single-goroutine use and provides no internal
// synchronization. If accessed from multiple goroutines, callers must
// synchronize all operations.
//
// # Time Semantics
//
// Create timestamps new records with time.Now(). If a specific time zone or
// clock source is required, handle that at a higher layer.
//
// # Usage
//
//	id := database.Create("Example", "Hello, world")
//	if log, ok := database.Get(id); ok {
//	    log.Text = "Updated"
//	    _ = database.Patch(id, log) // true if the record existed
//	}
//	database.Delete(id) // no-op if id is absent
//
// Package database.
package database
