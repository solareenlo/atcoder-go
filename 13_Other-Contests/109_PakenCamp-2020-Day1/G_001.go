package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	var l, r, x [400]int
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &l[i], &r[i], &x[i])
	}

	ans := -1
	for k := 0; k < (1 << N); k++ {
		var a [20]int
		for i := 1; i <= N; i++ {
			a[i] = a[i-1]
			if (k & (1 << (i - 1))) != 0 {
				a[i]++
			}
		}
		ok := true
		for i := 0; i < M; i++ {
			if a[r[i]]-a[l[i]-1] != x[i] {
				ok = false
				break
			}
		}
		if ok {
			ans = max(ans, popcount(uint32(k)))
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}
