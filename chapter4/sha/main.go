package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

var sha = flag.String("sha", "sha256", "hash that you want to use")

func main() {
	flag.Parse()
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		if *sha == "sha256" {
			c := sha256.Sum256([]byte(scan.Text()))
			fmt.Printf("%v\n", c)
		} else if *sha == "sha384" {
			c := sha512.Sum384([]byte(scan.Text()))
			fmt.Printf("%v\n", c)
		} else if *sha == "sha512" {
			c := sha512.Sum512([]byte(scan.Text()))
			fmt.Printf("%v\n", c)
		} else {
			log.Fatal("some")
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
}
