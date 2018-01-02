package main

import "fmt"

func a () func() string  {
	return func() string {
		return "Hello World"
	}

}

func main()  {

	b := a()
	fmt.Println(b())
}