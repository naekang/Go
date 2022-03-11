package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A, B, C, cnt int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &A, &B, &C)

	if B > C || B == C {
		cnt--
	} else {
		cnt = A/(C-B) + 1
	}

	fmt.Fprintf(writer, "%d\n", cnt)
}
