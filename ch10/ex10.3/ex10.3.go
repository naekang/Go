package main

import "fmt"

// if-else문을 좀 더 정리하기 위해 사용
func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("첫째 날")
		fmt.Println("오늘은 팀미팅이 있습니다.")
	case 2:
		fmt.Println("둘째 날")
		fmt.Println("새로운 팀원 면접이 있습니다.")
	case 3:
		fmt.Println("셋째 날")
		fmt.Println("설계안을 확정하는 날입니다.")
	case 4:
		fmt.Println("넷째 날")
		fmt.Println("예산을 확정하는 날입니다.")
	default:
		fmt.Println("프로젝트를 진행하세요.")
	}
}
