package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, file := range counts {
		total := 0
		for _, n := range file {
			total += n
		}
		if total <= 1 {
			continue
		}

		fmt.Printf("L:%s\nN:%d\tF:", line, total)
		for name := range file {
			fmt.Printf("%s ", name)
		}
		fmt.Println()
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "" {
			continue
		}
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][f.Name()]++
	}
	// ignoring potential errors from input.Error
}
