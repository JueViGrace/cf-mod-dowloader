package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	searchUrl *url.URL
	headers   map[string][]string = map[string][]string{
		"Accept":    {"application/json"},
		"x-api-key": {API_KEY},
	}
	modsResponse *APIResponse
)

func searchMods(mod ModFile, flags Flags) (*Mod, error) {
	fmt.Println("Initializing searchUrl")
	searchUrl, err := url.Parse(SEARCH_URL)
	if err != nil {
		return nil, err
	}

	fmt.Println("Initializing api response")
	modsResponse = new(APIResponse)

	fmt.Println("Initializing api response")
	modsResponse = new(APIResponse)

	fmt.Println("Adding query parameters to the request")
	params := url.Values{}
	params.Add("gameId", "432")
	params.Add("modLoaderType", "Forge")
	params.Add("gameVersion", flags.GameVersion)
	params.Add("filterText", mod.Name)

	fmt.Println("Encoding query parameters")
	searchUrl.RawQuery = params.Encode()

	fmt.Printf("Creating http request for url: %s\n", searchUrl.String())
	req, err := http.NewRequest("GET", searchUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Adding the headers")
	req.Header = headers

	fmt.Printf("Making the request for mod: %s\n", mod.Name)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("Reading body")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	fmt.Println("Unmarshaling body")
	err = json.Unmarshal(body, modsResponse)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
