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
		// Using the go key word, we can start a new go routine
		// By default Go will only attempt to use one core
		// We can override this and that means that Go can run more than
		// one routine at a time
		// Concurrency is not Parallelism
		// Concurrency - We can have multiple threads executing code.
		// If one thread blocks, another one is picked up and worked on
		// Parallelism - Multiple threads executed at the exact same time.
		// Requires multiple CPU's
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
