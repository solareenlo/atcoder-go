package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	N int
	B = [200005]int{}
)

func add(a, b int) {
	x := a
	for x <= N {
		B[x] += b
		x += x & -x
	}
}

func sum(a int) int {
	res := 0
	x := a
	for x > 0 {
		res += B[x]
		x -= x & -x
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var Q int
	fmt.Fscan(in, &N, &Q)

	for i := 0; i < Q; i++ {
		var t, k int
		fmt.Fscan(in, &t, &k)
		if t == 1 {
			s := sum(N) - sum(max(N-k, k-N-1))
			if s&1 != 0 {
				fmt.Fprintln(out, N*2+1-k)
			} else {
				fmt.Fprintln(out, k)
			}
		} else {
			add(k, 1)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
