package main

import "fmt"

// Scanf() 함수는 서식에 맞춘 입력을 받음
// func Scanf(format string, a ...interface{}) (n int, err error)
func main() {
	var a int
	var b int

	n, err := fmt.Scanf("%d %d\n", &a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
