package main

import (
	"os"

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
					mod.Name = c.FirstChild.Data
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
