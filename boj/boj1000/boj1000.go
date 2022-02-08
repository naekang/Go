package main

import "fmt"

func main() {
	var a int8
	var b int8

	fmt.Scanln(&a, &b)

	fmt.Println(a + b)
}
