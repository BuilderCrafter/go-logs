package main

import (
	"flag"
	"fmt"
	"logs_application/helpers"
	"logs_application/logs"
)

func main() {
	// Optional: allow preload via flag.
	var jsonPath string
	flag.StringVar(&jsonPath, "load", "spaceship_logs.json", "load logs from JSON on startup")
	flag.Parse()

	if jsonPath != "" {
		if err := logs.LoadLogsFromJSON(jsonPath); err != nil {
			fmt.Printf("Failed to load %q: %v", jsonPath, err)
			helpers.WaitForEnter("Continuing without preloaded logs. Press Enter…")
		}
	}
	for {
		helpers.ClearScreen()
		helpers.WelcomeMessage()
		logs.PrintLogs()
		fmt.Print("Enter log number to open, '+' to create new log, or 'q' to quit:")
		input := helpers.ReadLine()
		num, cmd := helpers.ParseMenuInput(input)
		switch {
		case num > 0:
			index := num - 1
			logs.OpenLog(uint8(index))
		case cmd == "q":
			fmt.Println()
			return

		case cmd == "+":
			logs.CreateLog()

		default:
			helpers.ClearScreen()
			helpers.WaitForEnter("Invalid input, please try again")
		}
	}
}
