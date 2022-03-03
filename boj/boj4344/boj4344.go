package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var C int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &C)
	defer writer.Flush()

	for i := 0; i < C; i++ {
		var N int
		fmt.Fscanf(reader, "%d ", &N)

		var scores = make([]float64, N)
		var sum, avg float64

		for j := 0; j < N; j++ {
			fmt.Fscanf(reader, "%f ", &scores[j])
			sum += scores[j]
		}
		avg = sum / float64(N)

		var proportion float64
		for _, val := range scores {
			if val > avg {
				proportion += (1 / float64(N))
			}
		}
		fmt.Fprintf(writer, "%.3f%s\n", math.Round(proportion*100000)/1000, "%")
	}
}
