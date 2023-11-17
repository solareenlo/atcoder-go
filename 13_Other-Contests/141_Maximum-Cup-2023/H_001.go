package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func f(a, m int) int {
	ret := 0
	for c := 1; c <= m; c++ {
		now := int(9e18)
		for i := 0; i < 61; i++ {
			if (a>>i)&1 == 0 {
				l := ((a >> i) | 1) << i
				bc := popcount(uint64(l))
				if bc <= c && c-bc <= i {
					l |= (1 << (c - bc)) - 1
					now = min(now, l)
				}
			}
		}
		ret = max(ret, now)
	}
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var m [50]int

	var N, X int
	fmt.Fscan(in, &N, &X)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &m[i])
	}
	a := X
	for i := 0; i < N; i++ {
		a = f(a, m[i])
	}
	fmt.Println(a)
}

func popcount(n uint64) int {
	return bits.OnesCount64(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
