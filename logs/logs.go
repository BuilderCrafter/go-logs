package logs

import (
	"fmt"
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
		fmt.Println("Invalid log index")
		fmt.Scanf("%v", input)
		return
	}
	log, ok := database.Get(log_index[index])
	if !ok {
		fmt.Println("LOG NOT FOUND")
		fmt.Scanf("%v", input)
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
		fmt.Scanf("%v", &input)
		input, val_input := helpers.LogParseInput(input)
		if !val_input {
			helpers.ClearScreen()
			fmt.Println("Invalid input")
			fmt.Scanf("%v", &input)
			continue
		}
		switch input {
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
	var input string
	helpers.ClearScreen()
	if index >= uint8(len(log_index)) {
		fmt.Println("Invalid log index")
		fmt.Scanf("%v", &input)
		return
	}
	log, ok := database.Get(log_index[index])
	if !ok {
		fmt.Println("LOG NOT FOUND")
		fmt.Scanf("%v", &input)
		return
	}
	for {
		helpers.ClearScreen()
		fmt.Println("Current log name:", log.Name)
		fmt.Print("Enter new name (max 50 characters/0 length to exit):")
		fmt.Scanf("%50s", &input)
		if len(input) > 50 {
			helpers.ClearScreen()
			fmt.Println("Input exceeds maximum length of 50 characters.")
			fmt.Scanf("%v", &input)
			continue
		} else if len(input) == 0 {
			helpers.ClearScreen()
			fmt.Println("Canceling edit operation.")
			fmt.Scanf("%v", &input)
			return
		}
		log.Name = input
		break // Exit the loop if a valid name is provided
	}
	helpers.ClearScreen()
	fmt.Printf("Current log text\n------------------------------ \n%v\n", helpers.WrapText(log.Text, 50))
	fmt.Println("Enter new text (0 length to exit)\n-------------------------------")
	fmt.Scanf("%255s", &input)
	if len(input) == 0 {
		helpers.ClearScreen()
		fmt.Println("Canceling edit operation.")
		fmt.Scanf("%v", &input)
		return
	}
	log.Text = input

	database.Patch(log_index[index], log)
	helpers.ClearScreen()
	fmt.Println("Log updated successfully")
	fmt.Scanf("%v", &input)
}

func Create_log() {
	var input string
	var input_name string
	var input_text string
	helpers.ClearScreen()
	if len(log_index) >= int(MAX_LENGTH) {
		helpers.ClearScreen()
		fmt.Println("Maximum number of logs reached. Cannot create new log.")
		fmt.Scanf("%v", &input)
		return
	}
	for {
		fmt.Println("Enter log name (max 50 characters/0 length to exit):")
		fmt.Scanf("%50s", &input)
		if len(input) > 50 {
			helpers.ClearScreen()
			fmt.Println("Input exceeds maximum length of 50 characters.")
			fmt.Scanf("%v", &input)
			return
		} else if len(input) == 0 {
			helpers.ClearScreen()
			fmt.Println("Canceling create operation.")
			fmt.Scanf("%v", &input)
			return
		}
		input_name = input
		break // Exit the loop if a valid name is provided
	}

	helpers.ClearScreen()
	fmt.Println("Enter new text (0 length to exit)\n-------------------------------")
	fmt.Scanf("%255s", &input)
	if len(input) == 0 {
		helpers.ClearScreen()
		fmt.Println("Canceling create operation.")
		fmt.Scanf("%v", &input)
		return
	}
	input_text = input

	log_id := database.Create(input_name, input_text)
	log_index = append(log_index, log_id)
	helpers.ClearScreen()
	fmt.Println("Log created successfully")
	fmt.Scanf("%v", &input)

}

func Delete_log(index uint8) {
	var input string
	helpers.ClearScreen()
	if index >= uint8(len(log_index)) {
		fmt.Println("Invalid log index")
		fmt.Scanf("%v", &input)
		return
	}
	log, ok := database.Get(log_index[index])
	if !ok {
		fmt.Println("LOG NOT FOUND")
		fmt.Scanf("%v", &input)
		return
	}
	database.Delete(log_index[index])
	log_index = append(log_index[:index], log_index[index+1:]...)
	helpers.ClearScreen()
	fmt.Println("Successfully deleted log: %v", log.Name)
	fmt.Scanf("%v", &input)
}
