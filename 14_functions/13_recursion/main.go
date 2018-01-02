package main

import (
	"fmt"

)

func factor (x int) int{

	if x == 0 {
		return 1
	}
	return x  * factor(x-1)
}



func main (){
	fmt.Println(factor(4))
}
