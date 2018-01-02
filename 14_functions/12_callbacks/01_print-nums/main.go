package main

import (
	"fmt"
)

func newFun (number []int, anc func(int)){
	for _, v := range number{
	anc(v)
	}
}


func main (){
	newFun ([]int{1,2,3,4,5}, func(n int) {
		fmt.Println(n)
	})
}