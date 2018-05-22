package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	sample1()
	sample2()
	sample3()
}

func sample1() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v\n", time.Now())
	// tmp := []interface{}{1, "apple", 1.4}
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %d, %s, %f\n", 1, "apple", 1.4)
}

func sample2() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
}

func sample3() {
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write(os.Stdout)
}
