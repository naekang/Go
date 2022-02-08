package main

import "fmt"

// %s: 기본 서식
// %q: 모든 특수문자가 기능을 잃고 그대로 출력
func main() {
	str := "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n"

	fmt.Println(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)
}
