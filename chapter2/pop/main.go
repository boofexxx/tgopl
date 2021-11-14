package main

import (
	"fmt"
	"tgopl/chapter2/popcount"
	"time"
)

func main() {
	var n uint64 = 77

	start := time.Now()
	fmt.Printf("Itera: %d ", popcount.PopCountIter(n))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Printf("Shift: %d ", popcount.PopCountShift(n))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Printf("Clear: %d ", popcount.PopCountClear(n))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Printf("Count: %d ", popcount.PopCount(n))
	fmt.Println(time.Since(start))
}
