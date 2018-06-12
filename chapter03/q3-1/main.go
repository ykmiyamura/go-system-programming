package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newfile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer newfile.Close()

	io.Copy(newfile, file)
}
