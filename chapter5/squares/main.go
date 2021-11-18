package main

import "fmt"

func square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := func() func() int {
		x := 0
		return func() int {
			x++
			return x * x
		}
	}()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
