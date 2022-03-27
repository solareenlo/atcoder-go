package main

import (
	"fmt"
	"math/bits"
)

func chk(x, y, c int) bool {
	if c%2 == 0 {
		x ^= y
	}
	if (x & (1 << (c / 2))) == 0 {
		return true
	}
	return false
}

func main() {
	var n, d1, d2 int
	fmt.Scan(&n, &d1, &d2)

	c1 := ctz(d1)
	c2 := ctz(d2)

	ct := 0
	for i := 0; i < (2*n-1)+1; i++ {
		for j := 0; j < (2*n-1)+1; j++ {
			if chk(i, j, c1) && chk(i, j, c2) {
				fmt.Println(i, j)
				ct++
				if ct == n*n {
					return
				}
			}
		}
	}
}

func ctz(x int) int {
	return bits.TrailingZeros64(uint64(x))
}
