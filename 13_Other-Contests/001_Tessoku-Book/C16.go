package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, K int
	fmt.Fscan(in, &N, &M, &K)
	que := make([]Tuple, 0)
	for i := 0; i < M; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		a--
		c--
		d += K
		que = append(que, Tuple{b, 1, a, i})
		que = append(que, Tuple{d, -1, c, i})
	}
	sortTUPLE(que)

	upd := make([]int, M)
	var dp [300005]int
	for _, x := range que {
		if x.b == -1 {
			dp[x.c] = max(dp[x.c], upd[x.d]+1)
		} else {
			upd[x.d] = dp[x.c]
		}
	}

	ans := -1
	for i := 0; i < N; i++ {
		ans = max(ans, dp[i])
	}
	fmt.Println(ans)
}

type Tuple struct {
	a, b, c, d int
}

func sortTUPLE(tup []Tuple) {
	sort.Slice(tup, func(i, j int) bool {
		if tup[i].a == tup[j].a {
			if tup[i].b == tup[j].b {
				if tup[i].c == tup[j].c {
					return tup[i].d < tup[j].d
				}
				return tup[i].c < tup[j].c
			}
			return tup[i].b < tup[j].b
		}
		return tup[i].a < tup[j].a
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
