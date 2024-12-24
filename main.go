package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	API_KEY string = os.Getenv("API_KEY")
	baseUrl *url.URL
	client  *http.Client = new(http.Client)
)

func main() {
	fmt.Println("Parsing base url")
	baseUrl, err := url.Parse("https://www.curseforge.com/api/v1/mods")
	if err != nil {
		fmt.Println("Error while parsing the base url")
		log.Fatal(err)
		return
	}

	fmt.Println("Parsing search url")
	searchUrl, err := url.Parse(baseUrl.String() + "/search")
	if err != nil {
		fmt.Println("Error pasing the search url")
		log.Fatal(err)
		return
	}

	fmt.Println("Getting flags from cmd")
	flags := pathFlag()

	fmt.Println("Reading html file with the mods")
	file, err := os.Open(flags.Path)
	if err != nil {
		fmt.Printf("Error while opening the file")
		log.Fatal(err)
		return
	}

	fmt.Println("Parsing html from file")
	mods, err := parseHtml(file)
	if err != nil {
		fmt.Println("Error parsing the file")
		log.Fatal(err)
		return
	}

	fmt.Println("Creating headers map")
	headers := map[string][]string{
		"Accept":    {"application/json"},
		"x-api-key": {API_KEY},
	}

	fmt.Println("Creating response mod list")
	modsList := make([]APIResponse, len(mods))

	fmt.Println("Iterating through local mods slice")
	for _, e := range mods {
		mod := new(APIResponse)

		fmt.Println("Adding query parameters to the request")
		params := url.Values{}
		params.Add("gameId", string(432))
		params.Add("gameVersion", flags.GameVersion)
		params.Add("searchFilter", e.Name)

		fmt.Println("Encoding query parameters")
		searchUrl.RawQuery = params.Encode()

		fmt.Println("Creating http request")
		req, err := http.NewRequest("GET", searchUrl.String(), nil)
		if err != nil {
			fmt.Println("Error while creating the request")
			log.Fatal(err)
			return
		}

		fmt.Println("Adding the headers")
		req.Header = headers

		fmt.Println("Making the request")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error while making the request")
			log.Fatal(err)
			return
		}
		defer resp.Body.Close()

		fmt.Println("Reading body")
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading the body")
			log.Fatal(err)
			return
		}
		if resp.StatusCode != 200 {
			log.Fatal(string(body))
			return
		}

		fmt.Println("Unmarshaling body")
		err = json.Unmarshal(body, mod)
		if err != nil {
			fmt.Println("Error while unmarshaling response body")
			return
		}

		fmt.Println("Adding response body to the new list")
		modsList = append(modsList, *mod)

		fmt.Println("Mod: ", string(body))
	}

	for _, e := range modsList {
		for _, d := range e.Data {
			fmt.Printf("New mod list: %v", d.Name)
		}
	}
}
