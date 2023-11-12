package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var K, M int
	fmt.Fscan(in, &K, &M)
	var repunits func(int, int) (int, int)
	repunits = func(y, x int) (int, int) {
		r, c, u, a := 0, 1, y, 10
		for x > 0 {
			if (x & 1) != 0 {
				r = (r + u) % M
				c = (c * a) % M
				u = (u * a) % M
			}
			u = (u * (a + 1)) % M
			a = (a * a) % M
			x /= 2
		}
		return r, c
	}
	ans := 0
	for i := 0; i < K; i++ {
		var c, d int
		fmt.Fscan(in, &c, &d)
		a, b := repunits(c, d)
		ans = ((ans * b % M) + a) % M
	}
	fmt.Println(ans)
}
