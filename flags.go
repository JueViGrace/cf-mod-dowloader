package main

import "flag"

type Flags struct {
	Path        string
	GameVersion string
}

func pathFlag() *Flags {
	flags := new(Flags)

	path := flag.String("path", "", "path of html file to read")
	gameVersion := flag.String("version", "1.20.1", "Game version to download")
	flag.Parse()

	flags.Path = *path
	flags.GameVersion = *gameVersion

	return flags
}
