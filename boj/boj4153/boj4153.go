package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		fmt.Fscanln(reader, &a, &b, &c)

		if a == 0 && b == 0 && c == 0 {
			break
		}

		if a*a+b*b == c*c {
			fmt.Fprintf(writer, "right\n")
		} else if a*a == b*b+c*c {
			fmt.Fprintf(writer, "right\n")
		} else if b*b == a*a+c*c {
			fmt.Fprintf(writer, "right\n")
		} else {
			fmt.Fprintf(writer, "wrong\n")
		}
	}
}
