package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	c := make([]int, 105)
	mx := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		c[a]++
		mx = max(mx, a)
	}

	fl := 1
	md := (mx + 1) / 2
	if c[md] != mx%2+1 {
		fl = 0
	}

	for i := md + 1; i <= mx; i++ {
		if c[i] < 2 {
			fl = 0
		}
	}
	if fl != 0 {
		fmt.Println("Possible")
	} else {
		fmt.Println("Impossible")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
