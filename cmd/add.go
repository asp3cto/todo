package cmd

import (
	"fmt"
	"github.com/asp3cto/todo/internal/db"
	io "github.com/asp3cto/todo/internal/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Run:   addHandler,
}

// Handler function for adding a todo
func addHandler(cmd *cobra.Command, args []string) {
	repo, err := db.NewTodoRepo(viper.GetString("todos_file"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Please, specify the name and the description of your todo 📝")

	name := io.GetInput("Name")
	description := io.GetInput("Description")

	if io.ConfirmTodo(name, description) {

		// Add new todo to the list
		todo := db.Todo{
			Name:        name,
			Description: description,
			Completed:   false,
		}
		_, err := repo.Insert(todo)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("Added todo ")
		color.Cyan(name + " 🎉🎉")
	} else {
		fmt.Println("Todo was not added")
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}
