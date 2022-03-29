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

	a, b := 0, 0
	for i := range s {
		if s[i]%2 != 0 {
			a += b
		} else {
			b++
		}
	}
	fmt.Println(a)
}
