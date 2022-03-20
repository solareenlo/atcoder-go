package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n int
	fmt.Fscan(in, &t, &n)

	c := make([]int, 500_500)
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		c[l]++
		c[r]--
	}

	for i := 0; i < t; i++ {
		c[i+1] += c[i]
		fmt.Fprint(out, c[i], " ")
	}
}
