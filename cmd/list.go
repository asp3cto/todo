package cmd

import (
	"fmt"
	"github.com/asp3cto/todo/internal/db"
	io "github.com/asp3cto/todo/internal/utils"
	"github.com/spf13/viper"
	"log"

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
	repo, err := db.NewTodoRepo(viper.GetString("todos_file"))
	if err != nil {
		log.Fatal(err)
	}
	switch {
	case all:
		todos, err := repo.SelectAll()
		if err != nil {
			log.Fatal(err)
		}
		if len(todos) == 0 {
			fmt.Println("You dont have any todos! Use `add` to create one :)")
			return
		}
		displayTodos(todos, verbose) // Display all todos
	case done:
		todos, err := repo.SelectByCompletedStatus(true)
		if err != nil {
			log.Fatal(err)
		}
		if len(todos) == 0 {
			fmt.Println("You dont have any done todos! Use `complete` to complete one :)")
			return
		}
		fmt.Printf("Done todos:\n\n")
		displayTodos(todos, verbose) // Display only completed todos
	default:
		todos, err := repo.SelectByCompletedStatus(false)
		if err != nil {
			log.Fatal(err)
		}
		if len(todos) == 0 {
			fmt.Println("You dont have any active todos! Use `add` to create one :)")
			return
		}
		fmt.Printf("Current todos:\n\n")
		displayTodos(todos, verbose) // Display only incomplete todos
	}
}

func displayTodos(todos []db.Todo, verbose bool) {
	for _, todo := range todos {
		io.VisualizeTodo(todo, verbose)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&all, "all", false, "Get all todos")
	listCmd.Flags().BoolVar(&done, "done", false, "Get only done todos")
	listCmd.MarkFlagsMutuallyExclusive("all", "done")

	listCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
}
