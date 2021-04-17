package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Call the count function to count the number of words
// received from the Standard Input and printing it out
func main() {
	fmt.Println(count(os.Stdin))
}

// Count the words received from the reader
func count(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	// Define the scanner split type to words (default is split by lines)
	scanner.Split(bufio.ScanWords)

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
