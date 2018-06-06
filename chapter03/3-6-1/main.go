package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `1行め
2行め
3行め`

func main() {
	reader()
	scanner()
}

func reader() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}

func scanner() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
