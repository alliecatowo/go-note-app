package main

import (
	"fmt"
	"example.com/go-note-app/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	
	if err != nil {
		fmt.Println("Error creating userNote", err)
		return
	}

	fmt.Println("Title:", title)
	fmt.Println("Content:", content)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	return value
}
