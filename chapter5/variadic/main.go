package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func min(n ...int) int {
	if len(n) == 0 {
		log.Fatal("min: passed 0 arguments")
	}
	min := n[0]
	for _, num := range n {
		if num < min {
			min = num
		}
	}
	return min
}

func max(n ...int) int {
	if len(n) == 0 {
		log.Fatal("max: passed 0 arguments")
	}
	max := n[0]
	for _, num := range n {
		if num > max {
			max = num
		}
	}
	return max
}

func stringsJoin(sep string, elems ...string) string {
	result := ""
	delim := ""
	for _, s := range elems {
		result += delim + s
		delim = sep
	}
	return result
}

func ElementsByTagName(n *html.Node, names ...string) []*html.Node {
	var res []*html.Node
	if n.Type == html.ElementNode {
		for _, name := range names {
			if n.Data == name {
				res = append(res, n)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, ElementsByTagName(c, names...)...)
	}
	return res
}

func main() {
	url := "https://gopl.io"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	images := ElementsByTagName(doc, "img")
	for _, img := range images {
		fmt.Printf("<%s", img.Data)
		for _, attr := range img.Attr {
			fmt.Printf(" %s=%s", attr.Key, attr.Val)
		}
		fmt.Printf(">\n")
	}
}
