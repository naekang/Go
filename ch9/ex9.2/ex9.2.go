package main

import "fmt"

func main() {
	tmp := 33

	if tmp > 28 {
		fmt.Println("에어컨을 켠다.")
	} else if tmp <= 3 {
		fmt.Println("히터를 켠다.")
	} else {
		fmt.Println("대기한다.")
	}
}
