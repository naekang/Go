package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n)

	dp := make([]int, 1001)

	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 10007
	}

	fmt.Fprintln(writer, dp[n])

}
