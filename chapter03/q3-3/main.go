package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("sample.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(writer, strings.NewReader("this is zipWriter sample."))

	_, err = zipWriter.Create("sample2.txt")
	if err != nil {
		panic(err)
	}
}
