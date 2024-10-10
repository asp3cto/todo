package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"todo/utils"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a todo",
	Run:   completeHandler,
}

func completeHandler(cmd *cobra.Command, args []string) {
	todos, err := utils.GetTodos()

	if err != nil {
		log.Fatal(err)
	}
	if len(todos) == 0 {
		fmt.Println("You dont have any active todos")
		return
	}

	for i, todo := range todos {
		if !todo.Completed {
			fmt.Printf("%d)", i+1)
			utils.VisualizeTodo(todo, false)
		}
	}
	todoID, err := strconv.Atoi(utils.GetInput("Give a number of a todo to complete"))
	if err != nil {
		log.Fatal(err)
	}
	if todoID < 1 || todoID > len(todos) {
		log.Fatal("Incorrect number!")
	}
	todos[todoID-1].Completed = true
	err = utils.SaveTodos(todos)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Completed todo!")
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
