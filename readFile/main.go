package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fn := os.Args[1]

	f, _ := os.Open(fn)

	fmt.Println("The contents of", fn, "is")
	io.Copy(os.Stdout, f)
}
