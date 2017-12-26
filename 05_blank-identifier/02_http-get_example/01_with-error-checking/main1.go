package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	a, _ := http.Get("https://google.com")
	b, _ := ioutil.ReadAll(a.Body)
	a.Body.Close()
	fmt.Printf("%s", b)
}
