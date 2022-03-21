package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)

	c := 63 - clz(uint64(n+1))
	t := c
	n -= (1 << c) - 1

	c += bits.OnesCount(uint(n))
	l := c
	fmt.Fprintln(out, l<<1)
	for i := 0; i < t; {
		if (n>>i)&1 != 0 {
			fmt.Fprint(out, c, " ")
			c--
		}
		i++
		fmt.Fprint(out, i, " ")
	}
	for i := 0; i < l; {
		i++
		fmt.Fprint(out, i, " ")
	}
}

func clz(x uint64) int {
	return bits.LeadingZeros64(x)
}
