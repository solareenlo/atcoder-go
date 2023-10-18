package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var g [30][]int
var x, y [30]int
var dx [4]int = [4]int{1, 0, -1, 0}
var dy [4]int = [4]int{0, 1, 0, -1}
var p3 [32]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a int
	fmt.Fscan(in, &a)
	var S, T [30]int
	for i := 0; i < a-1; i++ {
		var s string
		fmt.Fscan(in, &s)
		p := int(s[0] - 'A')
		fmt.Fscan(in, &s)
		q := int(s[0] - 'A')
		g[p] = append(g[p], q)
		g[q] = append(g[q], p)
		S[i] = p
		T[i] = q
	}
	p3[0] = 1
	for i := 1; i <= 31; i++ {
		p3[i] = p3[i-1] * 3
	}
	dfs(0, -1, -1, 0, 0, 0)
	zx := make([]int, a)
	zy := make([]int, a)
	for i := 0; i < a; i++ {
		zx[i] = x[i]
		zy[i] = y[i]
	}
	sort.Ints(zx)
	sort.Ints(zy)

	var px, py [32]int
	for i := 0; i < a; i++ {
		px[i] = lowerBound(zx, x[i]) * 3
		py[i] = lowerBound(zy, y[i]) * 3
	}

	var ret [120][120]byte
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			ret[i][j] = '.'
		}
	}
	for i := 0; i < a; i++ {
		ret[px[i]][py[i]] = byte('A' + i)
	}
	for i := 0; i < a-1; i++ {
		if px[S[i]] == px[T[i]] {
			for j := min(py[S[i]], py[T[i]]) + 1; j < max(py[S[i]], py[T[i]]); j++ {
				ret[px[S[i]]][j] = '-'
			}
		} else {
			for j := min(px[S[i]], px[T[i]]) + 1; j < max(px[S[i]], px[T[i]]); j++ {
				ret[j][py[S[i]]] = '|'
			}
		}
	}
	fmt.Fprintln(out, 100, 100)
	for i := 0; i < 100; i++ {
		fmt.Fprintln(out, string(ret[i][:100]))
	}
}

func dfs(a, b, c, X, Y, d int) {
	x[a] = X
	y[a] = Y
	t := 0
	for i := 0; i < len(g[a]); i++ {
		if b == g[a][i] {
			continue
		}
		if c == t {
			t++
		}
		dfs(g[a][i], a, t^2, X+p3[31-d]*dx[t], Y+p3[31-d]*dy[t], d+1)
		t++
	}
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

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
