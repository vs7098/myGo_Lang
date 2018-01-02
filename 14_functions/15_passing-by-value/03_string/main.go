package main

import "fmt"

func main() {

	name := "Vinay"
	fmt.Println(name) // Todd

	changeMe(name)

	fmt.Println(name) // Todd
}

func changeMe(z string) {
	fmt.Println(z) // Todd
	z = "Singh"
	fmt.Println(z) // Rocky
}
