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

	var q [80]int

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		var d, k, x int
		fmt.Fscan(in, &d, &k, &x)
		q[0] = 1
		res := int(1e18)
		for i := 1; i <= d; i++ {
			q[i] = q[i-1]*k + 1
		}
		for i := 0; i <= d; i++ {
			if q[i] >= x {
				cnt := 1
				if i == d {
					cnt = 0
				}
				p := q[i] - x
				for j := i - 1; j >= 0; j-- {
					cnt += p / q[j]
					p %= q[j]
				}
				res = min(res, cnt)
			}
		}
		fmt.Fprintln(out, res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
