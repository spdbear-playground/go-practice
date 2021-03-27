package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com/")
	fmt.Printf("%v\n%v", resp, err)
}
