package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"os"
)

const headerContentType = "Content-Type"

const ctHTML = "text/html"

func main() {
    for _, url := range os.Args[1:] {
        err := title(url)
        if err != nil {
        fmt.Printf("title: %s error: %v", url, err)
        }
    }
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get(headerContentType)
	if contentType != ctHTML && !strings.HasPrefix(contentType, ctHTML) {
		return fmt.Errorf("%s has type %s, not html", url, contentType)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("Parsing %s as HTML: %v", url, err)
	}

	visitNode := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" && node.FirstChild != nil {
			fmt.Println(node.FirstChild.Data)
		}
	}

	forEachNode(doc, visitNode, nil)

	return nil
}

func forEachNode(node *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(node)
	}
}
