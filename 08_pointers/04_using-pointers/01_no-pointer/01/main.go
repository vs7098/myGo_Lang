package main

import "fmt"

func zero(z int)  {

	z=0

}

func main()  {
	x := 10

	fmt.Println(&x)
	zero(x)


}