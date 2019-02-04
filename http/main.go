package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Make us a built in function of go that creates a variable of type with
	// a determined N value of of elements
	bs := make([]byte, 99999)

	// The read function does not expand the byte slice, so it is work
	// always creating a larger byte slice than is needed
	bw, err := resp.Body.Read(bs)

	fmt.Println(string(bs), bw)
}
