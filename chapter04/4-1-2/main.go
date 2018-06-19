package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()" + time.Now().String())

	go func() {
		fmt.Println("sub() is running" + time.Now().String())
		time.Sleep(time.Second)
		fmt.Println("sub() is finished" + time.Now().String())
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("end sub()" + time.Now().String())
}
