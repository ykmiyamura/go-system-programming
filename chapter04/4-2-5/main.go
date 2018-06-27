package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second)
		fmt.Println("sub() is finished.")
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("all tasks are finished.")
}
