package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	ch := time.After(10 * time.Second)
	<-ch
	fmt.Println("end")
}
