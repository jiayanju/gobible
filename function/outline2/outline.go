package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := "http://gopl.io"
	doc, err := GetWebDoc(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline2 : %v", err)
		os.Exit(1)
	}

	forEachNode(doc, startElement, endElement)
}

func GetWebDoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s:%s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return doc, nil
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

var depth int

func startElement(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", 2*depth, "", node.Data)
		depth++
	}
}

func endElement(node *html.Node) {
	if node.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", 2*depth, "", node.Data)
	}
}
