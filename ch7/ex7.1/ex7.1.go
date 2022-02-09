package main

import "fmt"

// func(함수키워드)
// Add(함수명)
// a int, b int(매개변수)
// int(반환 타입)
func Add(a int, b int) int {
	return a + b
}

// argument: 함수를 호출할 때 입력하는 값
// parameter: 함수가 외부로부터 입력받는 변수
func main() {
	c := Add(3, 6)
	fmt.Println(c)
}
