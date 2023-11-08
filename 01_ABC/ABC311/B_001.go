package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var as [101]bool

	var a, b int
	fmt.Fscan(in, &a, &b)
	for i := 1; i <= a; i++ {
		var c string
		fmt.Fscan(in, &c)
		for j := 0; j < b; j++ {
			if c[j] == 'x' {
				as[j] = true
			}
		}
	}
	d, e := 0, 0
	for j := 0; j < b; j++ {
		if as[j] {
			d = 0
		} else {
			d++
		}
		e = max(d, e)
	}
	fmt.Println(e)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
