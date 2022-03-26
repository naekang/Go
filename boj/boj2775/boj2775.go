package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, k, n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &k)
		fmt.Fscanln(reader, &n)
		fmt.Fprintln(writer, sol(k, n))
	}
}

func sol(k, n int) (cnt int) {
	if k == 1 {
		for i := 0; i < n; i++ {
			cnt += (i + 1)
		}
		return cnt
	}
	for i := 0; i < n; i++ {
		cnt += sol(k-1, i+1)
	}
	return cnt
}
