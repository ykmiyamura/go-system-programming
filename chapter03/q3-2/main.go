package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	// rReader := rand.Reader
	lReader := io.LimitReader(rand.Reader, 1024)
	file, err := os.Create("random.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, lReader)
}
