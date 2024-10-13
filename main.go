package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	command := flag.String("command", "", "Enter a command to run")

	flag.Parse()

	for {
		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()
	}

}

func runCommand(command string) {
	fmt.Println(command)
	switch command {
	case "add":
		addTask()
	case "update":
		updateTask()
	case "delete":
		deleteTask()
	case "mark-in-progress":
		markTask()
	case "mark-done":
		doneTask()
	case "list":
		listTask()
	case "exit":
		fmt.Println("Please enter the valid command")
		os.Exit(0)
	default:
		fmt.Println("Please enter the valid command")
		os.Exit(0)

	}

}

func addTask() {}

func updateTask() {}

func deleteTask() {}

func markTask() {}

func doneTask() {}

func listTask() {}
