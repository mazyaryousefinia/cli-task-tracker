package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

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

func getTasks() []Task {
	f, err := os.OpenFile("./tasks.json", os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	var tasks []Task

	err = decoder.Decode(&tasks)

	if err != nil {
		panic(err)
	}
	return tasks
}

func saveTasks(tasks []Task) {

	f, err := os.OpenFile("./tasks.json", os.O_RDWR|os.O_TRUNC, 0644)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(tasks)

	if err != nil {
		panic(err)
	}
}

func updateTaskById(taskId int, description string) ([]Task, error) {
	tasks := getTasks()

	for i := range tasks {
		if tasks[i].ID == taskId {
			tasks[i].Description = description
			now := time.Now().UTC()
			tasks[i].UpdatedAt = now
			return tasks, nil
		}
	}
	return tasks, errors.New("Cant find task")
}

func addTask() {
	var description string
	fmt.Println("Please enter task description")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	description = scanner.Text()

	tasks := []Task(getTasks())
	now := time.Now().UTC()
	task := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	tasks = append(tasks, task)
	saveTasks(tasks)
	fmt.Println("Task Created Successfully")

}

func updateTask() {
	var description string

	fmt.Println("Please enter tasks id")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	taskId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("error on entering task id")
	}

	fmt.Println("Please enter tasks new description")
	scanner.Scan()
	description = scanner.Text()
	tasks, err := updateTaskById(taskId, description)
	if err != nil {
		fmt.Println(err)
		return
	}
	saveTasks(tasks)
	fmt.Println("Task update successfully")

}

func deleteTask() {}

func markTask() {}

func doneTask() {}

func listTask() {}
