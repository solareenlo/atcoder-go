package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	a, b := 0, 0
	for _, c := range s {
		if c == '#' {
			a++
		} else if a > 0 {
			b++
			a--
		}
	}
	fmt.Println(b)
}
