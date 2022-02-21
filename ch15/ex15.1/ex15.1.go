package main

import "fmt"

func main() {
	// 큰따옴표로 묶으면 특수 문자가 동작함
	str1 := "Hello\t'World'\n"

	str2 := `Go is "awesome"!\nGo is simple and\t'powerful'`

	fmt.Println(str1)
	fmt.Println(str2)
}
