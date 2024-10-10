package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"todo/utils"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all th todos",
	Run:   clearHandler,
}

func clearHandler(cmd *cobra.Command, args []string) {
	color.Red("After this, all todos will be removed!")
	if utils.ConfirmClear() {
		err := utils.ClearTodos()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Cleared all todos")

	}

}

func init() {
	rootCmd.AddCommand(clearCmd)
}
