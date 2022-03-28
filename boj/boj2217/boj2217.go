package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var N int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)

	var rope = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &rope[i])
	}
	sort.Ints(rope)

	var ans int
	for i := len(rope) - 1; i >= 0; i-- {
		weight := rope[i] * (len(rope) - i)
		if weight > ans {
			ans = weight
		}
	}

	fmt.Fprintln(writer, ans)
}
