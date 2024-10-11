package utils

import (
	"fmt"
	"github.com/Songmu/prompter"
	"github.com/asp3cto/todo/internal/db"
	"github.com/fatih/color"
)

// Helper function to get user input with validation
func GetInput(prompt string) string {
	return prompter.Prompt(prompt, "")
}

// Helper function to confirm user input
func confirm(prompt string) bool {
	return prompter.YN(prompt, false)
}

// Function to confirm and display todo details
func ConfirmTodo(name, description string) bool {
	fmt.Println("Confirm your todo:")
	fmt.Printf("\t%s\n\t%s\n", name, description)
	return confirm("Is information right?")
}

func ConfirmClear() bool {
	return confirm("Clear all todos?")
}

func VisualizeTodo(todo db.Todo, verbose bool) {
	var titleStatus string
	if todo.Completed {
		titleStatus += color.GreenString(" ✔")
	} else {
		titleStatus += color.RedString(" ✘")
	}

	underlined := color.New(color.Underline).SprintFunc()

	fmt.Println(titleStatus, underlined(todo.Name))
	if verbose {
		fmt.Println("  " + todo.Description)
		fmt.Println()
	}
}
