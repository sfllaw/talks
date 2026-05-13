package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	// BEGIN OMIT
	r := strings.NewReader("Hello, 世界")

	b := make([]byte, 5)
	for {
		n, err := r.Read(b)     // Read len(b) bytes at a time // HL
		fmt.Printf("%s", b[:n]) // Process b before checking err
		if err != nil {
			if err == io.EOF { // io.EOF is a sentinel error
				break

			}
			log.Fatal(err)
		}
	}
	// END OMIT
}
