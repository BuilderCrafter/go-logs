package data

import "math/rand"

var database = make(map[string]string)

func Get(key string) string {
	value, exists := database[key]
	if exists {
		return value
	} else {
		return ""
	}
}

func Create(key string, value string) bool {
	// Simulate failiure to write
	randomNum := rand.Intn(10)
	if randomNum <= 8 {
		database[key] = value
		return true
	} else {
		return false
	}
}

func Patch(key string, value string) bool {
	_, exists := database[key]
	if exists {
		database[key] = value
		return true
	} else {
		return false
	}
}

func Delete(key string) bool {
	// Simulate failiure to delete
	randomNum := rand.Intn(10)
	if randomNum <= 8 {
		delete(database, key)
		return true
	} else {
		return false
	}
}
