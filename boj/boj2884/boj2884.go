package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var H, M int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &H, &M)

	if M >= 45 {
		fmt.Printf("%d %d", H, M-45)
	} else if H == 0 {
		fmt.Printf("%d %d", 23, 15+M)
	} else {
		fmt.Printf("%d %d", H-1, 15+M)
	}
}
