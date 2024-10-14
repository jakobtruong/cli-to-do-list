package app

import (
	"fmt"
)

type Task struct {
	Text      string
	Completed bool
}

// Prints out all available options to user
func showMenu(options []string) {
	fmt.Println("\nMenu:")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	fmt.Println(fmt.Printf("%d. Exit", len(options)+1))
}

func getUserInput() string {
	return ""
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show")
		return
	}
	fmt.Println("Tasks:")
	for i, task := range tasks {
		status := " "
		if task.Completed {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Text)
	}

}

func Run() {
	tasks := []Task{}
	options := []string{
		"Show Tasks",
		"Add Task",
		"Mark Task as Completed",
		"Save Tasks to File",
	}
	// while loop
	for {
		showMenu(options)
		userOption := getUserInput()
		switch userOption {
		case "1":
			showTasks(tasks)
		case "2":
			// Add Task
		case "3":
			// Mark Task as Complete
		case "4":
			// Save Tasks to csv file.
		case "5":
			fmt.Println("Exiting the To-Do application.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

	}
}
