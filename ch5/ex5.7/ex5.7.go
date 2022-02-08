package main

import "fmt"

// Scanln() 함수는 한 줄을 입력받아서 인수로 들어온 변수 메모리 주소에 값을 채움
// func Scanln(a ...interface{}) (n int, err error)
func main() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
