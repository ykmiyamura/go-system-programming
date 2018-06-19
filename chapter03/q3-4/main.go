package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "http.ResponseWriter sample")
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(writer, strings.NewReader("this is zipWriter sample."))
	w.Header().Set("Content-Type", "application/zip")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
