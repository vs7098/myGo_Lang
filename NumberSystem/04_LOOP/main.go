package main

import "fmt"
func main() {
	for i :=0 ; i <10 ; i++ {
		fmt.Printf("%d \t - %b \t  - %#X \n", i, i, i)
	}
}
