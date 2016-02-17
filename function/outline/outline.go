package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"os"

	"golang.org/x/net/html"
)

func main() {
	url := "https://golang.org"
	doc, err := html.Parse(strings.NewReader(GetWebContent(url)))
	if err != nil {
		fmt.Printf("outline: %v\n", err)
		os.Exit(1)
	}

	outline(nil, doc)

	fmt.Println("Print page by Breadth First Search")

	queue := []*html.Node{doc}
	outlineWidth(queue)
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

// note the usage of the slice which is not always changed.
func outline(stack []string, node *html.Node) {
	if node.Type == html.ElementNode {
		stack = append(stack, node.Data)
		fmt.Println(stack)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func outlineWidth(queue []*html.Node) {
	if len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.Data)

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode {
				queue = append(queue, c)
			}
		}

		outlineWidth(queue)
	}
}
