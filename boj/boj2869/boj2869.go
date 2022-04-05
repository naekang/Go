package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A, B, V int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &A, &B, &V)

	var ans = 1;

	if (V-A)%(A-B) == 0 {
		ans += (V - A) / (A - B)
	} else {
		ans += (V-A)/(A-B) + 1
	}

	fmt.Println(ans)
}
