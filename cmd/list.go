package cmd

import (
	"fmt"
	"log"
	"todo/models"
	"todo/utils"

	"github.com/spf13/cobra"
)

var all, done bool
var verbose bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of todos",
	Long: `Get a list of todos:
- if '--all' is set, prints all todos
- if '--done' is set, prints only done todos
- if '--verbose' is set, prints with description
- if no flags provided, prints only current todos`,
	Run: listHandler,
}

func listHandler(cmd *cobra.Command, args []string) {
	todos, err := utils.GetTodos()
	if err != nil {
		log.Fatal("Unable to get todos:", err)
	}

	if len(todos) == 0 {
		fmt.Println("You don't have any todos. Create some with the 'add' command. :)")
		return
	}

	switch {
	case all:
		fmt.Printf("Your todos:\n\n")
		displayTodos(todos, verbose, nil) // Display all todos
	case done:
		fmt.Printf("Done todos:\n\n")
		displayTodos(todos, verbose, func(todo models.Todo) bool {
			return todo.Completed
		}) // Display only completed todos
	default:
		fmt.Printf("Current todos:\n\n")
		displayTodos(todos, verbose, func(todo models.Todo) bool {
			return !todo.Completed
		}) // Display only incomplete todos
	}
}

func displayTodos(todos []models.Todo, verbose bool, filter func(todo models.Todo) bool) {
	for _, todo := range todos {
		if filter == nil || filter(todo) {
			utils.VisualizeTodo(todo, verbose)
		}
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&all, "all", false, "Get all todos")
	listCmd.Flags().BoolVar(&done, "done", false, "Get only done todos")
	listCmd.MarkFlagsMutuallyExclusive("all", "done")

	listCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
}
