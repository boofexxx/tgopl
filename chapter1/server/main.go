package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"tgopl/lissajous"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/counter/", counterHandler)
	http.HandleFunc("/lissajous", lissajousHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counterHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}

var (
	res     = 0.001
	cycles  = 5
	size    = 500
	nframes = 10
	delay   = 8
)

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		switch k {
		case "res":
			res, err = strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Print(err)
			}
		case "cycles":
			cycles, err = strconv.Atoi(v[0])
			if err != nil {
				log.Print(err)
			}
		case "size":
			size, err = strconv.Atoi(v[0])
			if err != nil {
				log.Print(err)
			}
		case "delay":
			delay, err = strconv.Atoi(v[0])
			if err != nil {
				log.Print(err)
			}
		case "nframes":
			nframes, err = strconv.Atoi(v[0])
			if err != nil {
				log.Print(err)
			}
		}
	}
	lissajous.Lissajous(w, res, cycles, size, nframes, delay)
}

