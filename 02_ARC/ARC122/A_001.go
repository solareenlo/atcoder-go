package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1e9 + 7

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, a int
	fmt.Fscan(in, &n, &a)

	b := 0
	c := 1
	d := 0
	for i := 0; i < n-1; i++ {
		var x int
		fmt.Fscan(in, &x)
		na := ((c+d)*x + a + b) % MOD
		nb := ((MOD-c)*x + a) % MOD
		nc := (c + d) % MOD
		nd := c
		a = na
		b = nb
		c = nc
		d = nd
	}

	fmt.Println((a + b) % MOD)
}
