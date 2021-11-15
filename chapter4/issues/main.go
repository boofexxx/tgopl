package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"tgopl/chapter4/github"
)

var period = flag.String("p", "all", "period of time when issues were created")

func main() {
	flag.Parse()
	result, err := github.SearchIssues(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	var from time.Time
	switch *period {
	case "month":
		from = now.AddDate(0, -1, 0)
	case "year":
		from = now.AddDate(-1, 0, 0)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	i := 0
	for _, item := range result.Items {
		if *period == "all" || item.CreatedAt.After(from) {
			fmt.Printf("%d %s ", i, item.State)
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
		i++
	}
}
