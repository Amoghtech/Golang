package main

import (
	"bufio"
	"example.com/notes/note"
	"example.com/notes/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}


func add[T int|float64|string](a, b T) T {
	return a+b;
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("1.5")
	title, content := getNoteData()
	text := getTodoData()

	todo, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(todo)
	if err != nil {
		return
	}
	printSomething(todo)

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputData(userNote)
}

func printSomething(value interface{}) {
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float64: ", value)
	// case string:
	// 	fmt.Println("String: ", value)
	// }

	// fmt.Println(value)

	intVal, ok := value.(int)
	if ok {
		// intVal += 1
		fmt.Println("Integer: ", intVal)
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float: ", floatVal)
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String: ", stringVal)
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed", err)
		return err
	}
	fmt.Println("Saving the note succeded")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note Content:")
	return title, content
}

func getTodoData() string {
	return getUserInput("Todo text: ")
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
