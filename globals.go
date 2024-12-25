package main

import (
	"net/http"
	"os"
)

var (
	API_KEY    string       = os.Getenv("API_KEY")
	BASE_URL   string       = os.Getenv("BASE_URL")
	SEARCH_URL string       = os.Getenv("SEARCH_URL")
	client     *http.Client = new(http.Client)
)
