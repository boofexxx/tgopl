package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	nletter, nspace, ndigit := 0, 0, 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		} else if unicode.IsLetter(r) {
			nletter++
		} else if unicode.IsSpace(r) {
			nspace++
		} else if unicode.IsDigit(r) {
			ndigit++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("number of letters: %d\n", nletter)
	fmt.Printf("number of digits: %d\n", ndigit)
	fmt.Printf("number of spaces: %d\n", nspace)
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%dinvalid UTF-8 characters\n", invalid)
	}
}
