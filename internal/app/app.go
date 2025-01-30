package app

import (
	"fmt"
	"strings"
)

type Task struct {
	Text      string
	Completed bool
}

// Returns menu showing all available options to user
// Using string.Builder here to efficiently build strings using write methods, minimizing memory copying.
func showMenu(options []string) string {
	var result strings.Builder
	result.WriteString("\nMenu:\n")
	for i, option := range options {
		fmt.Fprintf(&result, "%d. %s\n", i+1, option)
	}
	fmt.Fprintf(&result, "%d. Exit\n\n", len(options)+1)

	return result.String()
}

func getUserInput() int {
	var user_input int
	fmt.Print("Enter choice: ")
	fmt.Scanln(&user_input)
	return user_input
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

func (tasks *Task) addTasks(task Task) {

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
		print(showMenu(options))
		userOption := getUserInput()
		switch userOption {
		case 1:
			showTasks(tasks)
		case 2:
			// Add Task

		case 3:
			// Mark Task as Complete
		case 4:
			// Save Tasks to csv file.
		case 5:
			fmt.Println("Exiting the To-Do application.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

	}
}
