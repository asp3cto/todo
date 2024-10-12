package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/asp3cto/todo/internal/api"
	"github.com/asp3cto/todo/internal/db"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"time"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull issues as todos from github and gitlab",
	Run:   pullHandler,
}

func pullGithubIssues() ([]api.GitHubIssue, error) {
	ghToken := viper.GetString("github_token")
	client := http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest("GET", "https://api.github.com/issues?state=all", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+ghToken)
	req.Header.Add("Accept", "application/vnd.github+json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var issues []api.GitHubIssue
	err = json.Unmarshal(resBody, &issues)
	if err != nil {
		return nil, err
	}
	return issues, nil
}

func pullHandler(_ *cobra.Command, _ []string) {
	repo, err := db.NewTodoRepo(viper.GetString("todos_file"))
	if err != nil {
		log.Fatal(err)
	}

	issues, err := pullGithubIssues()
	if err != nil {
		log.Fatal(err)
	}

	var todos []db.Todo
	for _, issue := range issues {
		todo := db.Todo{
			Name:        issue.Title,
			Description: issue.Description,
			Completed:   issue.State != "open",
		}
		todos = append(todos, todo)
	}
	var processed int
	for _, todo := range todos {
		todoFromRepo, err := repo.GetByName(todo.Name)
		if err != nil {
			log.Fatal(err)
		}
		// we didnt find a todo in db and todo from api is not done - we add
		if todoFromRepo.Name == "" && !todo.Completed {
			_, err := repo.Insert(todo)
			if err != nil {
				log.Fatal(err)
			}
			processed++
			fmt.Print("Added todo ")
			color.Cyan(todo.Name + " 🎉🎉")

		} else if todoFromRepo.Name != "" && !todoFromRepo.Completed && todo.Completed {
			// this means we got a todo completion - we need to update it
			err := repo.CompleteByName(todo.Name)
			if err != nil {
				log.Fatal(err)
			}
			processed++
			fmt.Println("Completed todo " + color.CyanString(todo.Name) + " " + color.GreenString("✔"))
		}
	}
	if processed == 0 {
		fmt.Println("Your todos are already up-to-date!")
	}

}

func init() {
	rootCmd.AddCommand(pullCmd)
}
