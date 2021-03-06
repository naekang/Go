package main

import "fmt"

// 함수 선언부에 변수명까지 지정해주면, return문으로 해당 변수를 명시적으로 반환하지 않아도 값 반환 가능
func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return // 명시적으로 반환할 값을 지정하지 않은 return문
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)

	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
