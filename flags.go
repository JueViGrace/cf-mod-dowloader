package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Flags struct {
	Path        string
	GameVersion string
}

func pathFlag() *Flags {
	flags := new(Flags)

	defaultPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := flag.String("path", fmt.Sprintf("%s/%s", defaultPath, "modlist.html"), "path of html file to read")
	gameVersion := flag.String("version", "1.20.1", "Game version to download")
	flag.Parse()

	flags.Path = *path
	flags.GameVersion = *gameVersion

	return flags
}
