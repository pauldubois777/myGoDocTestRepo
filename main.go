package main

import (
	"fmt"
	"myGoDocTestRepo/myfirstpackage"
)

// main to be my first main for doc
func main() {
	result := myfirstpackage.GetHello("Paul")
	fmt.Println(result)
}
