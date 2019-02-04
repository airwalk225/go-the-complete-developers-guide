package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stakeoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, l := range links {
		go checkLink(l)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
	} else {
		fmt.Println(link, "is up and running")
	}
}
