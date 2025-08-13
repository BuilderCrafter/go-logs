// Package database provides an in-memory store for log records and the Log
// model type. It also includes a helper to import logs from a JSON file.
package database

import (
	"encoding/json"
	"os"
	"sync/atomic"
	"time"
)

// Log represents a log record.
// It includes a unique ID, a human-readable Name, the creation Date, and Text.
// The ID field is excluded from JSON serialization/deserialization.
type Log struct {
	ID   uint64    `json:"-"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
	Text string    `json:"text"`
}

// store is the in-memory map keyed by log ID. It emulates a database.
// Note: this package does not provide synchronization for concurrent access.
var store = make(map[uint64]Log)

// nextID holds the next unique ID issued for newly created logs.
// It is incremented atomically to avoid ID collisions.
var nextID uint64

// Get returns the Log with the given id and a boolean indicating whether
// such a log exists in the store.
func Get(id uint64) (Log, bool) {
	value, ok := store[id]
	return value, ok
}

// Create creates a new Log with the provided name and text, assigns it
// a unique ID and current timestamp, stores it, and returns the new ID.
func Create(name string, text string) uint64 {
	id := atomic.AddUint64(&nextID, 1)
	store[id] = Log{
		ID:   id,
		Name: name,
		Date: time.Now(),
		Text: text,
	}
	return id
}

// Patch updates the Log with the given id to the provided value.
// It returns true if the log existed and was updated, or false if no
// log with that id exists.
func Patch(id uint64, value Log) bool {
	_, exists := store[id]
	if exists {
		store[id] = value
		return true
	} else {
		return false
	}
}

// Delete removes the Log with the given id from the store.
// It is a no-op if the id does not exist.
func Delete(id uint64) {
	delete(store, id)
}

// LoadFromJSON reads logs from a JSON file at path and inserts them
// into the store, assigning fresh IDs. It returns the slice of new IDs
// in insertion order and any error encountered.
//
// The JSON is expected to represent an array of objects compatible with
// Log's exported fields (ID is ignored due to the json:"-" tag).
func LoadFromJSON(path string) ([]uint64, error) {
	var idIndex []uint64
	blob, err := os.ReadFile(path)
	if err != nil {
		return idIndex, err
	}
	var logs []Log
	if err := json.Unmarshal(blob, &logs); err != nil {
		return idIndex, err
	}
	for _, r := range logs {
		r.ID = atomic.AddUint64(&nextID, 1)
		idIndex = append(idIndex, r.ID)
		store[r.ID] = r
	}
	return idIndex, nil
}
