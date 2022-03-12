package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	fmt.Fprintf(writer, "%d\n", a*(b%10))
	fmt.Fprintf(writer, "%d\n", a*((b%100)/10))
	fmt.Fprintf(writer, "%d\n", a*(b/100))
	fmt.Fprintf(writer, "%d\n", a*b)
}
