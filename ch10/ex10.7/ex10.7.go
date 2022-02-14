package main

import "fmt"

// const 열거값에 따라 수행되는 로직을 변경할 때 switch문 주로 사용

type ColorType int // 별칭 ColorType을 선언하고 const 열거값 정의
const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}
