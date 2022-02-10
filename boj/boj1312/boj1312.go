package main

import "fmt"

func main() {
	var A, B, N int
	fmt.Scan(&A, &B, &N)

	A %= B

	for i := 1; i < N; i++ {
		A = A * 10 % B
	}

	fmt.Println(A * 10 / B)
}
