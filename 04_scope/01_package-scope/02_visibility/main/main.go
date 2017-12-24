package main

import (
	"fmt"
	"github.com/vs7098/myGo_Lang/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
