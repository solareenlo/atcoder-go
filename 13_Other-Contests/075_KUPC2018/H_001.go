package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 201000
	const mod = 1000000007

	var n, m, s int
	fmt.Fscan(in, &n, &m, &s)

	q := make([][][]int, N)
	for i := 0; i < m; i++ {
		var l, r, b int
		fmt.Fscan(in, &l, &r, &b)
		q[r] = append(q[r], []int{l, b, i})
	}

	pre := make([][]pair, N)
	for i := 1; i < s+1; i++ {
		pre[i] = append(pre[i], pair{0, 0})
	}

	var dp [N]int
	dp[0] = 1
	var mx [N]int
	for r := 1; r < n+1; r++ {
		sort.Slice(q[r], func(i, j int) bool {
			if q[r][i][0] == q[r][j][0] {
				if q[r][i][1] == q[r][j][1] {
					return q[r][i][2] < q[r][j][2]
				}
				return q[r][i][1] < q[r][j][1]
			}
			return q[r][i][0] < q[r][j][0]
		})
		q[r] = reverseOrderVec(q[r])
		dp[r] = dp[r-1] * s % mod
		for _, p := range q[r] {
			l := p[0]
			col := p[1]
			if mx[col] >= l {
				continue
			}
			z := -dp[l-1]
			it := lowerBoundPair(pre[col], pair{l, -1})
			if it < len(pre[col]) {
				z = (z - (pre[col][len(pre[col])-1].y - pre[col][it-1].y)) % mod
			}
			z %= mod
			if z < 0 {
				z += mod
			}
			pre[col] = append(pre[col], pair{r, (pre[col][len(pre[col])-1].y + z) % mod})
			mx[col] = l
			dp[r] = (dp[r] + z) % mod
		}
	}
	fmt.Println(dp[n])
}

type pair struct {
	x, y int
}

func lowerBoundPair(a []pair, x pair) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].x >= x.x
	})
	return idx
}

func reverseOrderVec(a [][]int) [][]int {
	n := len(a)
	res := make([][]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
