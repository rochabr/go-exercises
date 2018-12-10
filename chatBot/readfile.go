package main

import (
	"io"
	"log"
	"os"
)

func readFile(f string) {
	file, err := os.Open(f) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, file)
}
