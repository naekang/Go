package main

import "fmt"

// Scan() 함수는 변수들의 메모리 주소를 인수로 받음
// func Scan(a ...interface{}) (n int, err error)
func main() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
