package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, A, B int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &T)

	for i := 0; i < T; i++ {
		fmt.Fscanln(reader, &A, &B)
		fmt.Fprintf(writer, "Case #%d: %d + %d = %d\n", i+1, A, B, A+B)
	}
}
