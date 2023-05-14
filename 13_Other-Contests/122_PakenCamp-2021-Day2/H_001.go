package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(in, &m)
	nmin := 1
	nmax := 1000000000000000000
	for i := 0; i < m; i++ {
		var p, s int
		fmt.Fscan(in, &p, &s)
		k := 0
		var cntt [100]int
		cntt[0] = 1
		for cntt[k] <= s {
			k++
			cntt[k] = cntt[k-1]*p + 1
		}
		ans := make([]int, k)
		memo := 0
		for j := 0; j < k-1; j++ {
			ans[j] = (s - memo) / cntt[k-1-j]
			memo += ans[j]
		}
		ans[k-1] = s - memo
		nmin = max(nmin, ans[k-1]*p)
		nmax = min(nmax, ans[k-1]*p+p-1)
		if nmin > nmax {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(nmin)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
