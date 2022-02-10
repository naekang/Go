package main

import "fmt"

/**
 * iota 사용하기
 * 0부터 시작해서 자동으로 1씩 증가
 * 만약 첫번째 값과 같은 규칙이 적용된다면 타입과 iota 생략 가능
 * 1 << iota -> 1 << 0 = 1
 * 1 << iota -> 1 << 1 = 2
 * 1 << iota -> 1 << 2 = 4
 * 괄호를 벗어나면 초기화
 */

const Pig int = 0
const Cow int = 1
const Chicken int = 2

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...")
	}
}

func main() {
	PrintAnimal(0)
	PrintAnimal(Cow)
	PrintAnimal(2)
}
