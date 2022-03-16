package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, K, ans int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N, &K)

	ans = factorial(N) / (factorial(N-K) * factorial(K))

	fmt.Fprintf(writer, "%d\n", ans)
}

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		return factorial(num-1) * num
	}
}
