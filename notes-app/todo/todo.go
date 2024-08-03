package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text   string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, os.ModeAppend)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Value must not be empty")
	}

	return Todo{
		content,
	}, nil
}
