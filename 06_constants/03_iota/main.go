package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)
const (
	d = iota
	e = iota
	f = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
