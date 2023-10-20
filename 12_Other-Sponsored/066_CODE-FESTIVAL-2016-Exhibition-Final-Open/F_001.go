package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = math.MaxInt

	type seg struct {
		l, r, s int
	}

	var n int
	fmt.Fscan(in, &n)
	m := n >> 1
	s := make([]seg, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i].l, &s[i].r)
		s[i].s = s[i].l + s[i].r
	}
	tmp := s[1:]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].s > tmp[j].s
	})

	var f [2510][2]int
	for i := 1; i <= m; i++ {
		f[i][0] = INF
	}

	x := 0
	for i := 1; i <= n; i++ {
		x ^= 1
		for j := 0; j < m+1; j++ {
			f[j][x] = INF
		}
		for j := 0; j < i && j <= m; j++ {
			k := i - j - 1
			if j < m {
				f[j+1][x] = min(f[j+1][x], f[j][x^1]+s[i].r+s[i].s*j)
			}
			if k < m {
				f[j][x] = min(f[j][x], f[j][x^1]+s[i].l+s[i].s*k)
			}
			if (n & 1) != 0 {
				f[j][x] = min(f[j][x], f[j][x^1]+m*s[i].s)
			}
		}
	}

	fmt.Println(f[m][x])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
