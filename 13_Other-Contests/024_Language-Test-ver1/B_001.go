package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var x, s string
	fmt.Fscan(in, &x, &s)
	for i := 0; i < len(s); i++ {
		if s[i] != x[0] {
			fmt.Printf("%c", s[i])
		}
	}
	fmt.Println()
}
