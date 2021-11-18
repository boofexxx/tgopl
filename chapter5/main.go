package main

import "fmt"

func rec() (r string) {
	defer func() {
		switch r := recover(); r {
		case "some":
			r = "some"
		default:
			r = "something is wrong"
		}
	}()
	panic("some")
}

func main() {
	fmt.Println("vim-go")
}
