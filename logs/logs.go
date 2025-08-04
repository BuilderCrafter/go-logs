package logs

import (
	"fmt"
	"test_project/database"
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

func Read_log(index uint8) {
	var input string
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
	fmt.Printf("Log:  %v\n", log.Name)
	fmt.Printf("Date: %v\n", log.Date)
	fmt.Println("------------------------------------------")
	fmt.Printf("%v", wrap(log.Text, 50))
	fmt.Scanf("%v", input)
}
