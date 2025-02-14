package app

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Text      string
	Completed bool
}

var options = []string{
	"Show Tasks",
	"Add Task",
	"Mark Task as Completed",
	"Save Tasks to File",
}

// Returns menu showing all available options to user
// Using string.Builder here to efficiently build strings using write methods, minimizing memory copying.
func getMenu(options []string) string {
	var result strings.Builder
	result.WriteString("\nMenu:\n")
	for i, option := range options {
		fmt.Fprintf(&result, "%d. %s\n", i+1, option)
	}
	fmt.Fprintf(&result, "%d. Exit\n\n", len(options)+1)

	return result.String()
}

// Returns valid int if user_input is valid or -1 if not valid
func getOption(user_input string, options []string) int {
	i, err := strconv.Atoi(user_input)
	if err != nil {
		return -1
	}
	if i <= 0 || i > len(options) {
		return -1
	}
	return i
}

func showTasks(w io.Writer, tasks []Task) {
	if len(tasks) == 0 {
		fmt.Fprintln(w, "No tasks to show")
		return
	} else {
		fmt.Fprintln(w, "Tasks:")
		for i, task := range tasks {
			status := " "
			if task.Completed {
				status = "x"
			}
			fmt.Fprintf(w, "%d. [%s] %s\n", i+1, status, task.Text)
		}
	}
}

func (tasks *Task) addTasks(task Task) {

}

func Run() {
	tasks := []Task{}
	// while loop
	for {
		var userOption int = 0
		for {
			print(getMenu(options))
			var user_input string
			fmt.Print("Enter choice: ")
			fmt.Scanln(&user_input)
			userOption = getOption(user_input, options)
			if userOption != -1 {
				break
			}
		}
		switch userOption {
		case 1:
			showTasks(os.Stdout, tasks)
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
