package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to data storage application!")
	fmt.Println("This application allows the storage of text logs")
	var input int8
	for {
		fmt.Println("Choose one of the following actions")
		fmt.Println("1 - Read logs")
		fmt.Println("2 - Write log")
		fmt.Println("3 - Delete logs")
		fmt.Println("0 - Exit")
		fmt.Scanf("%v", &input)
		switch input {
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		case 1:
			fmt.Println("Reading to storage...")
			time.Sleep(1000000000)
			fmt.Println("Done!")
		case 2:
			fmt.Println("Writing to storage...")
			time.Sleep(1000000000)
			fmt.Println("Done!")
		default:
			fmt.Println("Invalid selection")
		}
	}
}
