package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	a       = [2003][2003]int{}
	b       = [2003][2003]int{}
	h, w, t int
	k       = [2003]float64{}
	d       = [2003]float64{}
	st      = [2003]int{}
)

func f(x float64, id int) float64 { return k[id]*x + d[id] }
func p(x, y int) float64          { return (d[x] - d[y]) / (k[y] - k[x]) }

type pair struct {
	x float64
	y int
}

func solve() {
	ans := -1000000
	t := 0
	for i := h; i > 0; i-- {
		k[i] = float64(i)
		d[i] = float64(a[h][i])
		for t > 1 && p(i, st[t-2]) < p(i, st[t-1]) {
			t--
		}
		st[t] = i
		t++
	}
	lw := make([]pair, 0)
	for i := 1; i <= w; i++ {
		lw = append(lw, pair{float64(b[w][i]) / float64(i), i})
	}
	sort.Slice(lw, func(i, j int) bool {
		return lw[i].x < lw[j].x
	})
	for _, x := range lw {
		for t > 1 && f(x.x, st[t-2]) > f(x.x, st[t-1]) {
			t--
		}
		aa := st[t-1]
		bb := x.y
		ans = max(ans, aa*b[w][bb]+bb*a[h][aa])
	}
	fmt.Println(ans)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 0; i < 2003; i++ {
		for j := 0; j < 2003; j++ {
			a[i][j] = -1 << 60
			b[i][j] = -1 << 60
		}
	}
	A := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &A[i])
		for j := i - 1; j > 0; j-- {
			a[i][j] = max(a[i][j], a[i-1][j])
		}
		cur := 0
		for j, k := 1, i; j <= i; {
			cur += A[k]
			a[i][j] = max(a[i][j], cur)
			j++
			k--
		}
	}
	B := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &B[i])
		for j := i - 1; j > 0; j-- {
			b[i][j] = max(b[i][j], b[i-1][j])
		}
		cur := 0
		for j, k := 1, i; j <= i; {
			cur += B[k]
			b[i][j] = max(b[i][j], cur)
			j++
			k--
		}
	}

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &h, &w)
		solve()
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
