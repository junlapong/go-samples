package main

import "fmt"

func main() {
	var i any
	i = 1
	i = "test"
	i = true

	fmt.Printf("%v\n", i)
}

