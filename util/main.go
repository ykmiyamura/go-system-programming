package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	fmt.Println(flag.Arg(0), flag.Arg(1))
	chapter := "../chapter" + flag.Arg(0)
	dirname := flag.Arg(1)

	os.Mkdir(chapter, 0755)
	os.Mkdir(chapter+"/"+dirname, 0755)
	os.Chdir(chapter + "/" + dirname)

	if !isExist("main.go") {
		f, err := os.Create("main.go")
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
