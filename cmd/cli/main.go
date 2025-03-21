package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/soohyeuk/cli_task_manager_go/pkg/models"
)

const taskFile = "tasks.json"

func loadTasks() ([]models.Task, error) {
	data, err := os.ReadFile(taskFile)
	if os.IsNotExist(err) {
		return []models.Task{}, nil
	}
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func saveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  task add <title>    - Add a new task")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "add" {
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Currently only 'add' command is implemented")
		os.Exit(1)
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Println("Please provide a task title")
		fmt.Println("Usage: task add <title>")
		os.Exit(1)
	}

	title := os.Args[2]
	newTask := models.Task{
		ID:        len(tasks) + 1,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	tasks = append(tasks, newTask)
	fmt.Printf("Added task: %s\n", title)

	if err := saveTasks(tasks); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
		os.Exit(1)
	}
}
