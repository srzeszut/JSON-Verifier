package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <json file>")
		return

	}
	jsonFilename := os.Args[1]
	fmt.Println("Verifying JSON file: ", jsonFilename)

	exampleJSON, err := os.ReadFile(jsonFilename)
	if err != nil {
		fmt.Println("Error reading JSON file: ", err)
		return
	}

	if verifyJSON(exampleJSON) {
		fmt.Println("JSON file is valid")
	} else {
		fmt.Println("JSON file is invalid")
	}

}
