package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	printer(doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "div" {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func populate(m map[string]int, n *html.Node) map[string]int {
	if m == nil {
		m = make(map[string]int)
	}
	if n.Type == html.ElementNode && n.Data == "div" {
		m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		populate(m, c)
	}

	return m
}

func printer(n *html.Node) {
	if n.Type == html.TextNode {
		if n.Parent.Type == html.ElementNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
			for _, s := range strings.Split(n.Data, "\n") {
				if len(s) != 0 {
					fmt.Println(s)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printer(c)
	}
}
