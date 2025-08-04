package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"test_project/logs"
)

func clearScreen() {
	// Clears the screen and moves the cursor to home
	ansiClear := "\033[2J\033[H"
	if _, err := fmt.Fprint(os.Stdout, ansiClear); err == nil {
		return
	}
}

func parseInput(s string) (num int, str string) {
	s = strings.TrimSpace(s)

	if n, err := strconv.Atoi(s); err == nil && n > 0 {
		return n, ""
	}
	return 0, s
}

func main() {
	var input string
	err := logs.LoadLogsFromJSON("spaceship_logs.json")
	if err != nil {
		panic(err)
	}
	for {
		clearScreen()
		fmt.Println("Welcome to logs storage application!")
		fmt.Println("This application allows the storage of text logs")
		fmt.Println("You can read, write, edit and delete logs")
		fmt.Println("Enter the number of log you would like to read/edit/delete")
		fmt.Println("Enter '+' if you would like to create new log")
		fmt.Println("---------------------LOGS------------------------")
		logs.Print_logs()
		fmt.Scanf("%v", &input)
		in_num, in_str := parseInput(input)
		if in_num != 0 {
			logs.Read_log(uint8(in_num) - 1)
		} else if in_str == "q" {
			os.Exit(0)
		}
		clearScreen()
	}
}
