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

	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var N, M int
		fmt.Fscan(in, &N, &M)
		ans := M
		cnt1 := make([]int, N)
		cnt2 := make([]int, N)
		for M > 0 {
			M--
			var u, v int
			fmt.Fscan(in, &u, &v)
			u--
			v--
			cnt1[u]++
			cnt2[v]++
		}
		for i := 0; i < N; i++ {
			ans -= min(cnt1[i], cnt2[i])
		}
		fmt.Fprintln(out, ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
