package main

import "fmt"

// 대입 연산자를 이용하여 복사
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	// 대입 연산자는 우변의 값을 좌변의 메모리 공간에 복사
	b = a

	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}
