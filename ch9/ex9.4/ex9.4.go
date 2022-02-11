package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

// && 연산은 좌변이 false이면 우변을 검사하지 않고 false 처리를 함 = 쇼트서킷(short-circuit)
func main() {
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 증가")
	}
	if true || IncreaseAndReturn() < 5 {
		fmt.Println("2 증가")
	}

	fmt.Println("cnt:", cnt)
}
