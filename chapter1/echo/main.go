package main

import (
	"fmt"
	"strings"
	"time"
)

func echoUsingJoin(args []string) string {
	return strings.Join(args[1:], " ")
}

func echoUsingIteration(args []string) string {
	s := ""
	sep := ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

var testData = []string{
	"something", "something", "something", "something",
	"something", "something", "something", "something",
	"something", "something", "something", "something",
	"something", "something", "something", "something",
}

func main() {
	fmt.Printf("echo using join:")
	start := time.Now()
	echoUsingJoin(testData)
	end := time.Now()
	fmt.Println(end.Sub(start))

	fmt.Printf("echo using iteartion:")
	start = time.Now()
	echoUsingIteration(testData)
	end = time.Now()
	fmt.Println(end.Sub(start))
}
