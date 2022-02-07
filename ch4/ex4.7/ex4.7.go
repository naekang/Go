package main

import "fmt"

// float32의 타입 제한으로 인해 큰 오차가 발생할 수 있기 때문에 실수 타입 사용 시 주의
func main() {
	var a float32 = 1234.123
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
