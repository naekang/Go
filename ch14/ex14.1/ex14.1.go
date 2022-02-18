package main

import "fmt"

func main() {
	var a int = 500
	var p *int // int 포인터 변수 p 선언

	p = &a

	fmt.Printf("p의 값 : %p\n", p)           // 메모리의 주소값 출력
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p) // p가 가리키는 메모리의 값 출력

	*p = 100 // p가 가리키는 메모리 공간의 값 변경
	fmt.Printf("a의 값: %d\n", a)
}
