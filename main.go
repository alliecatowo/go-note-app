package main

import (
	"fmt"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println("Error getting note data:", err)
		return
	}

	fmt.Println("Title:", title)
	fmt.Println("Content:", content)

}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title:")

	if err != nil {
		fmt.Println("Error getting title:", err)
		return "", "", err
	}

	content, nil := getUserInput("Note content:")

	if err != nil {
		fmt.Println("Error getting content:", err)
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	return value, nil
}
