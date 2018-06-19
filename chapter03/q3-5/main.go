package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("hello, i am gopher.")
	copyN(os.Stdout, reader, 10)
}

func copyN(dest io.Writer, src io.Reader, length int64) {
	r := io.LimitReader(src, length)
	io.Copy(dest, r)
}
