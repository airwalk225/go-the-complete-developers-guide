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

	// By default go will not wait for additional Go routines to finish
	// We have to use a Channel to keep the main the Go routine open,
	// or to pass information around between Go routines
	c := make(chan string)

	// To send messages through the channel
	// channel <- 5 - Send the value of 5 into this channel
	// myNumber <- channel - Wait for the value to be sent into the channel
	//                       When we get one, assign the value to myNumber
	// fmt.print(<- channel) - Wait for the value to be sent into the channel.
	//                         When we get one, log it out immediately

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
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
	} else {
		fmt.Println(link, "is up and running")
	}
}
