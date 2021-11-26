package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var mu sync.WaitGroup // not needed but just to be here for future idk =3
	input := bufio.NewScanner(c)
	ticker := time.NewTicker(10 * time.Second)
	done := make(chan struct{})
	go func() {
		for input.Scan() {
			ticker.Reset(10 * time.Second)
			mu.Add(1)
			go func() {
				echo(c, input.Text(), 1*time.Second)
				mu.Done()
			}()
		}
		done <- struct{}{}
	}()
	select {
	case <-ticker.C:
		break
	case <-done:
		break
	}
	mu.Wait() // it's not needed anymore. like we have 7 seconds after the last echo
	ticker.Stop()
	c.(*net.TCPConn).CloseWrite()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
