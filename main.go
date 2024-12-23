package main

import (
	"bytes"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	flags := pathFlag()

	file, err := os.Open(flags.Path)
	if err != nil {
		fmt.Printf("Opening file error: %s", err)
		return
	}

	mods, err := parseHtml(file)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	API_KEY := os.Getenv("API_KEY")
	headers := map[string][]string{
		"Accept":    {"application/json"},
		"x-api-key": {API_KEY},
	}

	for _, e := range mods {

		data := bytes.NewBuffer([]byte{})

		fmt.Printf("Mod: %v\n", e)
	}
}
