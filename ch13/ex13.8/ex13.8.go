package main

import (
	"fmt"
	"unsafe"
)

// 8바이트보다 작은 필드는 8바이트 크기를 고려해서 몰아서 배치
// 메모리 패딩 해결
type User struct {
	A int8
	B int8
	C int8
	D int
	E int
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
