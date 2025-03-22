package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func findTaskByID(tasks []models.Task, id int) (int, *models.Task) {
	taskMap := make(map[int]int) // map[taskID]sliceIndex
	for i, task := range tasks {
		taskMap[task.ID] = i
	}

	if idx, exists := taskMap[id]; exists {
		return idx, &tasks[idx]
	}
	return -1, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("There's no command provided. Please provide a command. If you need help, try 'task --help' or 'task -h'.")
		os.Exit(1)
	}

	command := os.Args[1]

	if command == "help" || command == "--help" || command == "-h" {
		fmt.Println("Usage:")
		fmt.Println("  task add <title>    - Add a new task")
		fmt.Println("  task list           - List all tasks")
		fmt.Println("  task done <id>      - Mark task as complete")
		fmt.Println("  task delete <id>    - Delete task")
		os.Exit(0)
	} else if command == "add" {
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

	} else if command == "list" {
		tasks, err := loadTasks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
			os.Exit(1)
		}

		if len(os.Args) > 2 {
			fmt.Println("Usage: task list")
			fmt.Println("Too many arguments provided. Please provide only 'list' command.")
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		for _, task := range tasks {
			status := "[ ]"
			if task.Completed {
				status = "[X]"
			}
			fmt.Printf("%s %d. %s (%s)\n", status, task.ID, task.Title, task.CreatedAt.Format("2006-01-02 15:04:05"))
		}
	} else if command == "done" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: task done <id>")
			os.Exit(1)
		}

		id := os.Args[2]
		idInt, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Invalid task ID")
			os.Exit(1)
		}

		tasks, err := loadTasks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
			os.Exit(1)
		}

		idx, task := findTaskByID(tasks, idInt)
		if task == nil {
			fmt.Printf("Task with ID %d not found\n", idInt)
			os.Exit(1)
		}

		tasks[idx].Completed = true
		fmt.Printf("Marked task %d as complete\n", idInt)

		if err := saveTasks(tasks); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			os.Exit(1)
		}
	} else if command == "delete" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: task delete <id>")
			os.Exit(1)
		}

		id := os.Args[2]
		idInt, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Invalid task ID")
			os.Exit(1)
		}

		tasks, err := loadTasks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
			os.Exit(1)
		}

		idx, task := findTaskByID(tasks, idInt)
		if task == nil {
			fmt.Printf("Task with ID %d not found\n", idInt)
			os.Exit(1)
		}

		tasks = append(tasks[:idx], tasks[idx+1:]...)
		fmt.Printf("Deleted task %d\n", idInt)

		if err := saveTasks(tasks); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			os.Exit(1)
		}

	} else {
		fmt.Println("Invalid command. Please try again.")
		os.Exit(1)
	}
}
