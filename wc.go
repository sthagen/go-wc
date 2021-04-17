package wc

import (
	"bufio"
	"io"
)

// Count the words received from the reader
func Count(r io.Reader, countLines bool) int {
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
