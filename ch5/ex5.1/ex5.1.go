package main

import "fmt"

// Print() 함수 입력값 출력
// Println() 함수 입력값 출력 + 개행
// Printf() format에 맞게 입력값 출력
func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print("a: ", a, "b: ", b)
	fmt.Println("a: ", a, "b: ", b, "f: ", f)
	fmt.Printf("a: %d, b: %d, f: %f\n", a, b, f)
}
