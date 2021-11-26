package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("commencing countdown. Press return to abort.")
	ticker := time.NewTicker(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-ticker.C:
			// do nothing
		case <-abort:
			fmt.Println("launch aborted!")
			return
		}
	}
	ticker.Stop()
	launch()
}

func launch() {
	fmt.Println("launched succesfully")
}
