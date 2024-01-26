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

	const N = 200200
	const M = 840

	var n, X, Y int
	fmt.Fscan(in, &n, &X, &Y)

	var p, t [N]int
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &p[i], &t[i])
	}

	ans := make([]int, 1010)
	ans2 := make([]int, 1010)
	for i := 0; i < M; i++ {
		ans[i] = 0
	}
	for i := n - 1; i > 0; i-- {
		ans2[0] = t[i] + ans[t[i]%M]
		for j := M - 1; j > 0; j-- {
			if j%p[i] != 0 {
				ans2[j] = ans2[(j+1)%M] + 1
			} else {
				ans2[j] = t[i] + ans[(t[i]+j)%M]
			}
		}
		for j := 0; j < M; j++ {
			ans[j] = ans2[j]
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 1; i <= q; i++ {
		var st int
		fmt.Fscan(in, &st)
		st += X
		fmt.Fprintln(out, Y+st+ans[st%M])
	}
}
