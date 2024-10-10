package utils

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"todo/models"
)

func InitStorage(path string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	return err
}

func GetTodos() ([]models.Todo, error) {
	content, err := os.ReadFile(viper.GetString("todos_file"))
	if err != nil && !os.IsNotExist(err) { // Only fail if the error is not due to the file not existing
		return nil, err
	}

	var todos []models.Todo
	if len(content) > 0 { // Unmarshal only if there is content
		err = json.Unmarshal(content, &todos)
		if err != nil {
			return nil, err
		}
	}
	return todos, nil
}

func SaveTodos(todos []models.Todo) error {
	// Marshal the list back into JSON
	b, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	// Open the file for writing and truncate (clear the file before writing new data)
	f, err := os.OpenFile(viper.GetString("todos_file"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Error while closing storage: ", err)
		}
	}(f)

	// Write the JSON to the file
	_, err = f.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func ClearTodos() error {
	f, err := os.OpenFile(viper.GetString("todos_file"), os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("could not open file for truncation: %v", err)
	}
	if err = f.Close(); err != nil {
		return fmt.Errorf("could not close file handler for after truncation: %v", err)
	}
	return nil
}
