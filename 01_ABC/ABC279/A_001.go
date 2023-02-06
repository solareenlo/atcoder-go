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
	x := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'w' {
			x += 2
		} else {
			x++
		}
	}
	fmt.Println(x)
}
