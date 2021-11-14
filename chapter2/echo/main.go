package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	for i := 0; i < 3; i++ {
		c := 5
		fmt.Printf("%p\n", &i)
		fmt.Printf("%p\n", &c)
	}
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
