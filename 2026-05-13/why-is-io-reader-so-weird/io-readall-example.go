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

	b, err := io.ReadAll(r) // HL
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)
	// END OMIT
}
