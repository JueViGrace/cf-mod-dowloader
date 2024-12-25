package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var (
	searchUrl *url.URL
	headers   map[string][]string = map[string][]string{
		"Accept":    {"application/json"},
		"x-api-key": {API_KEY},
	}
	modsResponse *APIResponse
)

func searchMods(modFile ModFile, flags Flags) (*Mod, error) {
	fmt.Println("Initializing searchUrl")
	searchUrl, err := url.Parse(SEARCH_URL)
	if err != nil {
		return nil, err
	}

	fmt.Println("Initializing api response")
	modsResponse = new(APIResponse)

	fmt.Println("Initializing api response")
	mod := new(Mod)

	fmt.Println("Adding query parameters to the request")
	params := url.Values{}
	params.Add("gameId", "432")
	params.Add("modLoaderType", "Forge")
	params.Add("gameVersion", flags.GameVersion)
	params.Add("filterText", modFile.Name)

	fmt.Println("Encoding query parameters")
	searchUrl.RawQuery = params.Encode()

	fmt.Printf("Creating http request for url: %s\n", searchUrl.String())
	req, err := http.NewRequest("GET", searchUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Adding the headers")
	req.Header = headers

	fmt.Printf("Making the request for modFile: %s\n", modFile.Name)
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

	for _, e := range modsResponse.Data {
		if strings.Contains(e.Name, strings.Split(modFile.Name, " ")[0]) {
			mod.ID = e.ID
			mod.Name = e.Name
			mod.LocalName = modFile.Name
			mod.LatestFiles = e.LatestFiles
		}
	}

	return mod, nil
}
