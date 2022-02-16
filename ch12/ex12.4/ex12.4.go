package main

import "fmt"

// range 키워드를 사용하여 인덱스와 요소값 두 개 다 반환 가능하다
func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i, v := range t {
		fmt.Println(i, v)
	}
}
