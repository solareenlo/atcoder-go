package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const maxn = 200010
	const INF = math.MaxInt - 1

	type pair struct {
		x, y int
	}

	var n, x int
	fmt.Fscan(in, &n, &x)
	var a, b [maxn]int
	sum := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		if a[i] == 1 {
			sum += b[i]
			i--
			n--
		}
	}
	if n == 0 {
		fmt.Println(1, sum)
		return
	}
	type edge struct {
		to, val int
	}
	G := make([][]edge, maxn)
	for i := 1; i <= n; i++ {
		v := make([]int, 0)
		for j := i; j <= n; j++ {
			v = append(v, j)
			for p := len(v) - 1; p > 0; p-- {
				if (a[v[p-1]]-1)*b[v[p]] >= (a[v[p]]-1)*b[v[p-1]] {
					break
				}
				v[p-1], v[p] = v[p], v[p-1]
			}
			val := 0
			for _, p := range v {
				val = val*a[p] + b[p]
			}
			if val > x {
				break
			}
			G[i] = append(G[i], edge{j + 1, val})
		}
	}
	var calc func(int) pair
	calc = func(p int) pair {
		dp := make([]pair, maxn)
		dp[1] = pair{sum, 0}
		for i := 2; i <= n+1; i++ {
			dp[i] = pair{INF >> 1, -1}
		}
		for i := 1; i <= n; i++ {
			for _, e := range G[i] {
				val := dp[i].x + e.val + p
				if dp[e.to].x > val || (dp[e.to].x == val && dp[e.to].y < dp[i].y+1) {
					dp[e.to] = pair{val, dp[i].y + 1}
				}
			}
		}
		return pair{dp[n+1].x - dp[n+1].y*p, dp[n+1].y}
	}
	l, r := 0, x
	var p int
	for l <= r {
		mid := (l + r) >> 1
		dp := calc(mid)
		if dp.x <= x {
			p = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	dp := calc(p)
	D := dp.y - (x-dp.x)/p
	M := dp.x + p*(dp.y-D)
	fmt.Println(D, M)
}
