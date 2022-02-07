package main

import "fmt"

func main() {
	var a int16 = 3456
	var c int8 = int8(a) // int16에서 int8로 변환 시 상위 1바이트가 사라짐

	fmt.Println(a)
	fmt.Println(c)
}
