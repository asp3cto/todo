package cmd

import (
	"fmt"
	"github.com/asp3cto/todo/internal/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init an app and create all the necessary files",
	Run:   initHandler,
}

func initHandler(cmd *cobra.Command, args []string) {
	_, err := db.NewTodoRepo(viper.GetString("todos_file"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully initialized app! 🚀🚀🚀")
}

func init() {
	rootCmd.AddCommand(initCmd)
}
