// Package logs provides a CLI-oriented, in-memory index of log IDs and
// helpers for listing, opening, editing, creating, and deleting logs.
package logs

import (
	"fmt"
	"logs_application/database"
	"logs_application/helpers"
	"strings"
)

// log_index holds the ordered list of log IDs as shown in the CLI.
// The order here (not the database) defines the menu numbering the user sees.
var log_index []uint64

// MAX_LENGTH is the maximum number of logs tracked by the CLI index.
const MAX_LENGTH uint8 = 255

// LoadLogsFromJSON loads log IDs from a JSON file and replaces the current index.
// It returns any error encountered during loading.
//
// The JSON loading/parsing is delegated to the database package.
func LoadLogsFromJSON(path string) error {
	id_index, err := database.LoadFromJSON(path)
	log_index = id_index
	return err
}

// Print_logs prints the 1-based menu of logs by name.
// Missing IDs (not present in the database) are reported as "LOG NOT FOUND".
func PrintLogs() {
	for index, id := range log_index {
		log, ok := database.Get(id)
		if ok {
			fmt.Printf("%3d - %v\n", index+1, log.Name) // 1-based numbering for UX
		} else {
			fmt.Printf("%3d - LOG NOT FOUND\n", index)
		}
	}
}

// Open_log starts an interactive view loop for the log at the given index.
// The user may edit (E), delete (X), or return (R) to the previous menu.
//
// If the index is invalid or the log is missing, an informational prompt is shown.
func OpenLog(index uint8) {
	var input string
	helpers.ClearScreen()

	// Validate that the menu index exists in the local index.
	if index >= uint8(len(log_index)) {
		helpers.WaitForEnter("Invalid log index")
		return
	}
	log, ok := database.Get(log_index[index])
	if !ok {
		helpers.WaitForEnter("LOG NOT FOUND")
		return
	}
	for {
		helpers.ClearScreen()
		fmt.Printf("Log:  %v\n", log.Name)
		fmt.Printf("Date: %v\n", log.Date)
		fmt.Println("------------------------------------------")
		fmt.Printf("%v\n", helpers.WrapText(log.Text, 50))
		fmt.Println("Edit: E | Delete: X | Return: R")
		fmt.Print("Input: ")
		input = helpers.ReadLine()
		input, val_input := helpers.LogParseInput(input)
		if !val_input {
			helpers.ClearScreen()
			helpers.WaitForEnter("Invalid input")
			continue
		}
		switch strings.ToLower(input) {
		case "x":
			DeleteLog(index)
			return
		case "e":
			EditLog(index)
			// Re-fetch updated log because Edit_log may have changed it.
			log, _ = database.Get(log_index[index])
			continue
		case "r":
			return
		}
	}
}

// Edit_log prompts the user to change the log's Name and Text at a given index.
// Name input is limited to 50 characters; sending an empty input cancels.
func EditLog(index uint8) {
	var input_name string
	var input_text string
	helpers.ClearScreen()
	if index >= uint8(len(log_index)) {
		helpers.WaitForEnter("Invalid log index")
		return
	}
	log, ok := database.Get(log_index[index])
	if !ok {
		helpers.WaitForEnter("LOG NOT FOUND")
		return
	}

	// Name edit loop with validation and cancel path.
	for {
		helpers.ClearScreen()
		fmt.Println("Current log name:", log.Name)
		fmt.Print("Enter new name (max 50 characters/0 length to exit):")
		input_name = helpers.ReadLine()
		if len(input_name) > 50 {
			helpers.ClearScreen()
			helpers.WaitForEnter("Input exceeds maximum length of 50 characters.")
			continue
		} else if len(input_name) == 0 {
			helpers.ClearScreen()
			helpers.WaitForEnter("Canceling edit operation.")
			return
		}
		log.Name = input_name
		break // Valid name provided
	}
	helpers.ClearScreen()
	fmt.Printf("Current log text\n------------------------------ \n%v\n", helpers.WrapText(log.Text, 50))
	fmt.Println("Enter new text (0 length to exit)\n-------------------------------")
	input_text = helpers.ReadLine()
	if len(input_text) == 0 {
		helpers.ClearScreen()
		helpers.WaitForEnter("Canceling edit operation.")
		return
	}
	fmt.Println(input_text)
	log.Text = input_text

	database.Patch(log_index[index], log)
	helpers.ClearScreen()
	helpers.WaitForEnter("Log updated successfully")
}

// Create_log interactively creates a new log, appending its ID to the CLI index.
// It enforces MAX_LENGTH on the number of tracked logs and allows canceling
// the operation with an empty input.
func CreateLog() {
	var input_name string
	var input_text string
	helpers.ClearScreen()
	if len(log_index) >= int(MAX_LENGTH) {
		helpers.ClearScreen()
		helpers.WaitForEnter("Maximum number of logs reached. Cannot create new log.")
		return
	}

	// Name input with length limit and cancel option.
	for {
		helpers.ClearScreen()
		fmt.Println("Enter log name (max 50 characters/0 length to exit):")
		input_name = helpers.ReadLine()
		if len(input_name) > 50 {
			helpers.ClearScreen()
			helpers.WaitForEnter("Input exceeds maximum length of 50 characters.")
			continue
		} else if len(input_name) == 0 {
			helpers.ClearScreen()
			helpers.WaitForEnter("Canceling create operation.")
			return
		}
		break // Valid name provided
	}

	helpers.ClearScreen()
	fmt.Println("Enter new text (0 length to exit)\n-------------------------------")
	input_text = helpers.ReadLine()
	if len(input_text) == 0 {
		helpers.ClearScreen()
		helpers.WaitForEnter("Canceling create operation.")
		return
	}

	log_id := database.Create(input_name, input_text)
	log_index = append(log_index, log_id)
	helpers.ClearScreen()
	helpers.WaitForEnter("Log created successfully")

}

// Delete_log removes the log at the given index from both the database and the CLI index.
// If the index is invalid or the log is missing, it shows an informational prompt.
func DeleteLog(index uint8) {
	helpers.ClearScreen()
	if index >= uint8(len(log_index)) {
		helpers.WaitForEnter("Invalid log index")
		return
	}
	log, ok := database.Get(log_index[index])
	if !ok {
		helpers.WaitForEnter("LOG NOT FOUND")
		return
	}
	database.Delete(log_index[index])

	// Remove the element at 'index' from log_index while preserving order.
	log_index = append(log_index[:index], log_index[index+1:]...)
	helpers.ClearScreen()
	fmt.Printf("Successfully deleted log: %v", log.Name)
	helpers.WaitForEnter(" ")
}
