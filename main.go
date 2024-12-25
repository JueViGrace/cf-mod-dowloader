package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("Getting flags from cmd")
	flags := pathFlag()

	fmt.Println("Reading html file with the mods")
	file, err := os.Open(flags.Path)
	if err != nil {
		fmt.Printf("Error while opening the file")
		log.Fatal(err)
		return
	}
	defer file.Close()

	fmt.Println("Parsing html from file")
	mods, err := parseHtml(file)
	if err != nil {
		fmt.Println("Error parsing the file")
		log.Fatal(err)
		return
	}

	modsList := make([]Mod, 0, 0)
	for _, mod := range mods {
		newMod, err := searchMods(mod, *flags)
		if err != nil {
			fmt.Println("Error searching mods")
			log.Fatal(err)
			break
		}
		modsList = append(modsList, *newMod)
	}

	for _, e := range modsList {
		fmt.Printf("Mods to download list: %v\n", e)
	}
}
