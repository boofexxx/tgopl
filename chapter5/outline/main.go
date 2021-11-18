package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("use: outline [url]")
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
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

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth = 0

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			fmt.Printf("%*s</%s", depth*2, "", n.Data)
		} else {
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
		}
		for _, a := range n.Attr {
			fmt.Printf(" %s=%s", a.Key, a.Val)
		}
		fmt.Printf(">\n")
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
