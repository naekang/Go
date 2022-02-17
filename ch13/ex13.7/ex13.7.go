package main

import (
	"fmt"
	"unsafe"
)

// 메모리 정렬을 위해 필드 사이에 공간을 띄우는 것을 메모리 패딩 이라고 함
type User struct {
	A int8
	B int
	C int8
	D int
	E int8
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}

// User 구조체는 1바이트 필드 3개와 8바이트 필드 2개로 구성 = 19바이트 차지
// 실제 구조는 메모리 패딩 때문에 40바이트 <- 1바이트 변수 A, C, E 모두에 7바이트씩 패딩됬기 때문
