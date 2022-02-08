package main

import "fmt"

// float64 표현방식으로 인해 다음과 같은 문제가 생길 수 있음
func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	fmt.Println(a + b)
}
