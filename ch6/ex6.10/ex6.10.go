package main

import "fmt"

// 대입 연산자
func main() {
	var a int = 10
	var b int = 20

	a, b = b, a

	fmt.Println(a, b)
}
