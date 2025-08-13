package logs

import (
	"fmt"
	"strings"
	"test_project/database"
	"test_project/helpers"
)

var log_index []uint64

const MAX_LENGTH uint8 = 255

func Num_of_logs() uint8 {
	return uint8(len(log_index))
}

func LoadLogsFromJSON(path string) error {
	id_index, err := database.InitFromJSON(path)
	log_index = id_index
	return err
}

func Print_logs() {
	for index, id := range log_index {
		log, ok := database.Get(id)
		if ok {
			fmt.Printf("%3d - %v\n", index+1, log.Name)
		} else {
			fmt.Println("%3d - LOG NOT FOUND\n", index)
		}
	}
}

func Open_log(index uint8) {
	var input string
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
			Delete_log(index)
			return
		case "e":
			Edit_log(index)
			log, _ = database.Get(log_index[index])
			continue
		case "r":
			return
		}
	}
}

func Edit_log(index uint8) {
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
		break // Exit the loop if a valid name is provided
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

func Create_log() {
	var input_name string
	var input_text string
	helpers.ClearScreen()
	if len(log_index) >= int(MAX_LENGTH) {
		helpers.ClearScreen()
		helpers.WaitForEnter("Maximum number of logs reached. Cannot create new log.")
		return
	}
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
		break // Exit the loop if a valid name is provided
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

func Delete_log(index uint8) {
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
	log_index = append(log_index[:index], log_index[index+1:]...)
	helpers.ClearScreen()
	fmt.Printf("Successfully deleted log: %v", log.Name)
	helpers.WaitForEnter(" ")
}
