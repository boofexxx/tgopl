package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var address = flag.String("address", "localhost:8080", "address")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *address)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Print("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.(*net.TCPConn).CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
