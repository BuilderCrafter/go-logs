package main

import (
	"fmt"
	"os"
	"test_project/helpers"
	"test_project/logs"
)

func main() {
	var input string
	err := logs.LoadLogsFromJSON("spaceship_logs.json")
	if err != nil {
		panic(err)
	}
	for {
		helpers.ClearScreen()
		helpers.WelcomeMessage()
		logs.Print_logs()
		fmt.Print("Enter log number to open, '+' to create new log, or 'q' to quit:")
		input = helpers.ReadLine()
		in_num, in_str := helpers.MenuParseInput(input)
		if in_num != 0 {
			logs.Open_log(uint8(in_num) - 1)
		} else if in_str == "q" {
			os.Exit(0)
		} else if in_str == "+" {
			logs.Create_log()
		} else {
			helpers.ClearScreen()
			helpers.WaitForEnter("Invalid input, please try again")
		}
	}
}
