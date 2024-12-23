package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	path := pathFlag()

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Opening file error: %s", err)
		panic(err)
	}

	links, err := parseHtml(file)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

	fmt.Println("Links: ", links)
}

func pathFlag() string {
	path := flag.String("path", "", "path of html file to read")
	flag.Parse()
	return *path
}

func parseHtml(file *os.File) ([]string, error) {
	doc, err := html.Parse(file)
	if err != nil {
		return nil, err
	}

	links := extractLinks(doc)

	return links, nil
}

func extractLinks(n *html.Node) []string {
	links := make([]string, 0, 0)
	var crawler func(*html.Node)
	crawler = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "li" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Data == "a" {
					for _, a := range c.Attr {
						if a.Key == "href" {
							fmt.Println("Link: ", a.Val)
							links = append(links, a.Val)
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawler(c)
		}
	}
	crawler(n)
	return links
}
