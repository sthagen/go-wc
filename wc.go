package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// Call the count function to count the number of words
// received from the Standard Input and printing it out
func main() {
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}

// Count the words received from the reader
func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)
	// Define the scanner split type to words (default is split by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
