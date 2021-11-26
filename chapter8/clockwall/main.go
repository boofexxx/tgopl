package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var newYork = flag.String("NewYork", "localhost:8010", "adress of NewYork")
var tokyo = flag.String("Tokyo", "localhost:8020", "adress of Tokyo")
var london = flag.String("London", "localhost:8030", "adress of London")

func main() {
	flag.Parse()
	connNewYork, err := net.Dial("tcp", *newYork)
	if err != nil {
		log.Fatal(err)
	}
	connTokyo, err := net.Dial("tcp", *tokyo)
	if err != nil {
		log.Fatal(err)
	}
	connLondon, err := net.Dial("tcp", *london)
	if err != nil {
		log.Fatal(err)
	}
	go mustCopy(os.Stdout, connNewYork)
	go mustCopy(os.Stdout, connTokyo)
	mustCopy(os.Stdout, connLondon)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
