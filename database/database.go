package database

import (
	"encoding/json"
	"os"
	"sync/atomic"
	"time"
)

type Log struct {
	ID   uint64    `json:"-"`
	Name string    `json:"name"`
	Date time.Time `json:"time"`
	Text string    `json:"text"`
}

var database = make(map[uint64]Log)
var nextID uint64

func Get(id uint64) (Log, bool) {
	value, ok := database[id]
	return value, ok
}

func Create(name string, text string) uint64 {
	id := atomic.AddUint64(&nextID, 1)
	database[id] = Log{
		ID:   id,
		Name: name,
		Date: time.Now(),
		Text: text,
	}
	return id
}

func Patch(id uint64, value Log) bool {
	_, exists := database[id]
	if exists {
		database[id] = value
		return true
	} else {
		return false
	}
}

func Delete(key uint64) {
	delete(database, key)
}

func InitFromJSON(path string) ([]uint64, error) {
	var id_index []uint64
	blob, err := os.ReadFile(path)
	if err != nil {
		return id_index, err
	}
	var logs []Log
	if err := json.Unmarshal(blob, &logs); err != nil {
		return id_index, err
	}
	for _, r := range logs {
		r.ID = atomic.AddUint64(&nextID, 1)
		id_index = append(id_index, r.ID)
		database[r.ID] = r
	}
	return id_index, nil
}
