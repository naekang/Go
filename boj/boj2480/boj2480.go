package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, c int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &a, &b, &c)

	if a == b && b == c && c == a {
		fmt.Fprintf(writer, "%d\n", 10000+(a*1000))
	} else if a == b || a == c {
		fmt.Fprintf(writer, "%d\n", 1000+(a*100))
	} else if b == c {
		fmt.Fprintf(writer, "%d\n", 1000+(b*100))
	} else {
		fmt.Fprintf(writer, "%d\n", int(math.Max(float64(a), math.Max(float64(b), float64(c)))*100))
	}

}
