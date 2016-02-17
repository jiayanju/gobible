package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "http://golang.org"
	content := GetWebContent(url)
	// 	fmt.Printf("%s", content)

	links := findLinks(content)
	fmt.Println("Links")
	for _, link := range links {
		fmt.Println(link)
	}
}

func GetWebContent(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("fetch: %v\n", err)
		return ""
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("fetch: reading %s:%v\n", url, err)
		return ""
	}

	return string(b)
}

func findLinks(webContent string) []string {
	doc, err := html.Parse(strings.NewReader(webContent))
	if err != nil {
		fmt.Errorf("%s", err)
		return nil
	}

	return visit(nil, doc)
}

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		fmt.Println("Find a")
		for _, a := range node.Attr {
			if a.Key == "href" {
				fmt.Printf("Find Link: %s\n", a.Val)
				links = append(links, a.Val)
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}
