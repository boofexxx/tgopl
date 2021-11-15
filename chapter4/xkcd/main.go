package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

const URL = "https://xkcd.com/"

var num = flag.Int("num", 1, "the number of page")

func main() {
	flag.Parse()
	q := url.QueryEscape(strconv.Itoa(*num) + "/info.0.json")
	resp, err := http.Get(URL + q)
	if err != nil {
		log.Fatalf("xkcd: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("xkcd: %v", resp.Status)
	}
	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		resp.Body.Close()
		log.Fatalf("xkcd: %v\n", err)
	}
	fmt.Printf("%s\n", comic.Transcript)
	fmt.Printf("%s\n", comic.Img)
}
