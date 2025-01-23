package main

import "fmt"

func smm(a, b int) int {
	return a + b
}

func main() {

	var a = 5
	var b = 6
	c := smm(a, b)
	fmt.Println("welcome to the Go world !! ", c)
}
