package cmd

import (
	"fmt"
	"github.com/asp3cto/todo/internal/db"
	io "github.com/asp3cto/todo/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a todo",
	Run:   completeHandler,
}

func completeHandler(_ *cobra.Command, _ []string) {
	repo, err := db.NewTodoRepo(viper.GetString("todos_file"))
	if err != nil {
		log.Fatal(err)
	}
	todos, err := repo.SelectByCompletedStatus(false)
	if err != nil {
		log.Fatal(err)
	}
	if len(todos) == 0 {
		fmt.Println("You dont have any active todos")
		return
	}

	indexToPK := make(map[int]int)
	for i, todo := range todos {
		indexToPK[i+1] = todo.ID
	}

	for i, todo := range todos {
		fmt.Printf("%d)", i+1)
		io.VisualizeTodo(todo, false)
	}
	todoID, err := strconv.Atoi(io.GetInput("Give a number of a todo to complete"))
	if err != nil {
		log.Fatal(err)
	}
	if todoID < 1 || todoID > len(todos) {
		log.Fatal("Incorrect number!")
	}
	err = repo.Complete(indexToPK[todoID])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Completed todo!")
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
