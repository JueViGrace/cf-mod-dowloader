package main

import (
	"os"
	"strings"

	"golang.org/x/net/html"
)

type ModFile struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func parseHtml(file *os.File) ([]ModFile, error) {
	doc, err := html.Parse(file)
	if err != nil {
		return nil, err
	}

	links := extractLinks(doc)

	return links, nil
}

func extractLinks(n *html.Node) []ModFile {
	mods := make([]ModFile, 0, 0)

	var crawler func(*html.Node)
	crawler = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "li" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				mod := new(ModFile)
				if c.Data == "a" {
					var modName string

					switch {
					case strings.Contains(c.FirstChild.Data, "["):
						modName = sliptModName(c.FirstChild.Data, "[")
					case strings.Contains(c.FirstChild.Data, "("):
						modName = sliptModName(c.FirstChild.Data, "(")
					case strings.Contains(c.FirstChild.Data, "/"):
						modName = sliptModName(c.FirstChild.Data, "/")
					}

					mod.Name = modName
					for _, a := range c.Attr {
						if a.Key == "href" {
							mod.Link = a.Val
						}
					}
				}
				mods = append(mods, *mod)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawler(c)
		}
	}
	crawler(n)
	return mods
}

func sliptModName(s string, separator string) string {
	var name string
	words := make([]string, 0, 0)
	words = strings.Split(s, " ")

	switch {
	case strings.Contains(words[0], separator):
		name = strings.Split(strings.Join(words[1:], " "), separator)[0]
	default:
		name = strings.Split(s, separator)[0]
	}

	return name
}
