package main

import (
	"bufio"
	"fmt"
	"os"
)

type query struct{ i, l int }

var (
	que = make([][]query, 1<<18)
	VW  = [1 << 18][2]int{}
	dp  = [10][100001]int{}
	dp2 = [1 << 9][2]int{}
	ans = [100000]int{}
)

func DFS(p, i int) {
	if i == 18 {
		return
	}
	v := VW[p][0]
	w := VW[p][1]
	if i < 9 {
		dp[i+1][0] = 0
		for j := 1; j <= 100000; j++ {
			dp[i+1][j] = max(dp[i+1][j-1], dp[i][j])
			if j >= w {
				dp[i+1][j] = max(dp[i+1][j], dp[i][j-w]+v)
			}
		}
		for _, q := range que[p] {
			ans[q.i] = dp[i+1][q.l]
		}
	} else {
		ii := i - 9
		for j := 0; j < 1<<ii; j++ {
			dp2[j+(1<<ii)][0] = dp2[j][0] + v
			dp2[j+(1<<ii)][1] = dp2[j][1] + w
		}
		for _, q := range que[p] {
			for j := 0; j < 1<<(ii+1); j++ {
				if dp2[j][1] <= q.l {
					ans[q.i] = max(ans[q.i], dp[9][q.l-dp2[j][1]]+dp2[j][0])
				}
			}
		}
	}
	DFS((p<<1)+1, i+1)
	DFS((p<<1)+2, i+1)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &VW[i][0], &VW[i][1])
	}

	var q int
	fmt.Fscan(in, &q)

	for i := 0; i < q; i++ {
		var v, l int
		fmt.Fscan(in, &v, &l)
		v--
		que[v] = append(que[v], query{i, l})
	}

	DFS(0, 0)

	for i := 0; i < q; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
