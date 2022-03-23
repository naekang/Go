package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	fmt.Fprintf(writer, "%d", Factorial(N))
}

func Factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * Factorial(num-1)
}
