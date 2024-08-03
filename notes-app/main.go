package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving the note succeded")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note Content:")
	return title, content
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
