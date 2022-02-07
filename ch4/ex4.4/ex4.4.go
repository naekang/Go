package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	var c int = int(b)
	d := float64(a * c)

	var e int64 = 7
	f := int64(d) * e

	var g int = int(b * 3) // 실수 타입에서 정수 타입으로 변환 시 소수점 이하는 반올림되지 않고 사라짐
	var h int = int(b) * 3

	fmt.Println(g, h, f)
}
