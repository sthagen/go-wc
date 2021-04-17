package main

import (
	"flag"
	"fmt"
	"os"

	wc "github.com/sthagen/go-wc"
)

// Call the count function to count the number of words
// received from the Standard Input and printing it out
func main() {
	lines := flag.Bool("l", false, "Count lines not words")
	flag.Parse()
	fmt.Println(wc.Count(os.Stdin, *lines))
}
