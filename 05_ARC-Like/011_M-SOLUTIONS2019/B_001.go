package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)
	n := len(s)

	c := 15
	for i := n - 1; i >= 0; i-- {
		if s[i] == 'x' {
			c--
		}
	}
	if c >= 8 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
