package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const APIURL = "http://www.omdbapi.com/"
const APIKey = "&apikey=3df0d9d2"

type Movie struct {
	Title    string
	Year     string
	Rated    string
	Released string
	Runetime string
	Genre    string
	Director string
	Writer   string
	Actors   string
	Plot     string
	Language string
	Country  string
	Awars    string
	Poster   string
}

func getMovie(name string) (*Movie, error) {
	resp, err := http.Get(APIURL + "?t=" + name + "&plot=full" + APIKey)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &movie, nil
}

func getPoster(movie *Movie) error {
	resp, err := http.Get(movie.Poster)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("search query failed: %s", resp.Status)
	}
	f, err := os.Create(movie.Title + ".jpeg")
	if err != nil {
		return err
	}
	io.Copy(f, resp.Body)

	return nil
}

func main() {
	name := flag.String("n", "", "the name of the movie")
	if len(os.Args) < 2 {
		flag.Usage()
		return
	}
	flag.Parse()

	movie, err := getMovie(*name)
	if err != nil {
		log.Fatal(err)
	}
	if err := getPoster(movie); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title: %s\nPlot: %s\n", movie.Title, movie.Plot)
}
