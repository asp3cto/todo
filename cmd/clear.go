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

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all th todos",
	Run:   clearHandler,
}

func clearHandler(_ *cobra.Command, _ []string) {
	color.Red("After this, all todos will be removed!")
	if io.ConfirmClear() {
		repo, err := db.NewTodoRepo(viper.GetString("todos_file"))
		if err != nil {
			log.Fatal(err)
		}
		err = repo.Clear()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Cleared all todos")

	}

}

func init() {
	rootCmd.AddCommand(clearCmd)
}
