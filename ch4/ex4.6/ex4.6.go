package main

import "fmt"

var g int = 10

// 변수는 중괄호 {} 범위를 벗어나면 사라짐
// g 는 전역변수
// m은 지역변수로 main() 함수 전체 사용 가능
// s도 지역변수 이지만 두번째 { } 내부에서만 사용 가능
func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	//m = s + 20 // Error
}
