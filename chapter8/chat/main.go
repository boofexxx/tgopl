package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected client
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// client outgoing message channels
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

var users []string

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	fmt.Fprintln(conn, "Active users:")
	for _, user := range users {
		fmt.Fprintf(conn, "%s ", user)
	}
	fmt.Fprintln(conn)
	users = append(users, who)
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	ticker := time.NewTicker(time.Minute * 5)
	input := bufio.NewScanner(conn)
	done := make(chan struct{})

	go func() {
		for input.Scan() {
			ticker.Reset(time.Minute * 5)
			messages <- who + ": " + input.Text()
		}
		done <- struct{}{}
	}()
	// NOTE: ignoring potential errors from input.Err()

	select {
	case <-done:
		break
	case <-ticker.C:
		ticker.Stop()
		break
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
