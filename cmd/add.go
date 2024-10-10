package cmd

import (
	"fmt"
	"github.com/asp3cto/todo/models"
	"github.com/asp3cto/todo/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
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
	fmt.Println("Please, specify the name and the description of your todo 📝")

	name := utils.GetInput("Name")
	description := utils.GetInput("Description")

	if utils.ConfirmTodo(name, description) {
		// Get list of Todo from storage
		todos, err := utils.GetTodos()
		if err != nil {
			log.Fatal("Error while getting todos from storage:", err)
		}

		// Add new todo to the list
		todo := models.Todo{
			Name:        name,
			Description: description,
			Completed:   false,
		}
		todos = append(todos, todo)

		err = utils.SaveTodos(todos)
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
