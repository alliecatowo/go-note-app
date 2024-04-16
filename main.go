package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-note-app/note"
	"example.com/go-note-app/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error creating userNote", err)
		return
	}

	userNote.Display()

	err = userNote.Save()

	if err != nil {
		fmt.Println("Error saving userNote", err)
		return
	}

	fmt.Println("Note saved successfully")

	todo.Display()
	err = todo.Save()
	if err != nil {
		fmt.Println("Error saving todo", err)
		return
	}
	fmt.Println("Todo saved successfully")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
