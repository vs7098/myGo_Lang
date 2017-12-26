package main

import "fmt"

const (
	_ = iota
	a = iota * 10
	b = iota * 10
)

func main() {
	fmt.Println(a)
	fmt.Println(b)

}
