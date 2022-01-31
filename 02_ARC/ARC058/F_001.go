package main

import (
	"bufio"
	"fmt"
	"os"
)

type pii struct{ x, y int }

const maxn = 2005
const maxk = 10005

var (
	n, k int
	z    = make([]int, maxk)
	p    = make([]int, maxk)
	st   [maxk]pii
	s    = make([]string, maxn)
	t    = make([]string, maxn)
	ff   [maxn][maxk]bool
	ok   [maxn][maxk]bool
)

func z_func(s string, z []int) {
	N := len(s)
	z[0] = N
	for i, l, r := 1, 0, 0; i < N; i++ {
		if i <= r && z[i-l] < r-i+1 {
			z[i] = z[i-l]
		} else {
			z[i] = max(0, r-i+1)
			for i+z[i] < N && s[i+z[i]] == s[z[i]] {
				z[i]++
			}
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
}

func exkmp(s, t string, z, p []int) {
	z_func(s, z)
	N := len(s)
	m := len(t)
	now := 0
	for now < min(N, m) && s[now] == t[now] {
		now++
	}
	p[0] = now
	for i, l, r := 1, 0, 0; i < m; i++ {
		if i <= r && z[i-l] < r-i+1 {
			p[i] = z[i-l]
		} else {
			p[i] = max(0, r-i+1)
			for i+p[i] < m && p[i] < N && t[i+p[i]] == s[p[i]] {
				p[i]++
			}
		}
		if i+p[i]-1 > r {
			l = i
			r = i + p[i] - 1
		}
	}
}

func comp(i int, a, b pii) int {
	fl := 1
	res := 0
	if a.x < b.x {
		a.x, b.x = b.x, a.x
		a.y, b.y = b.y, a.y
		fl *= -1
	}
	x := b.x
	y := a.x
	if x+p[x] < y && p[x] < b.y {
		if t[i-1][x+p[x]] < s[i][p[x]] {
			res = -1
		} else {
			res = 1
		}
	} else if z[y-x] < a.y && y-x+z[y-x] < b.y {
		if s[i][z[y-x]] < s[i][y-x+z[y-x]] {
			res = -1
		} else {
			res = 1
		}
	}
	return res * fl
}

func solve() {
	t[1] = s[1]
	ff[1][0] = true
	ff[1][len(s[1])] = true
	for i := 2; i <= n; i++ {
		l := len(s[i])
		for i := range z {
			z[i] = 0
			p[i] = 0
		}
		exkmp(s[i], t[i-1], z, p)
		top := 0
		for j := 0; j <= k; j++ {
			if ok[i+1][k-j] {
				var cur pii
				a := pii{j - l, l}
				b := pii{j, 0}
				if j >= l && ff[i-1][j-l] && ff[i-1][j] {
					if comp(i, a, b) == -1 {
						cur = a
					} else {
						cur = b
					}
				} else if j >= l && ff[i-1][j-l] {
					cur = a
				} else if ff[i-1][j] {
					cur = b
				} else {
					continue
				}
				for ; top > 0 && comp(i, cur, st[top]) == -1; top-- {
					ff[i][st[top].x+st[top].y] = false
				}
				if top == 0 || comp(i, cur, st[top]) == 0 {
					top++
					st[top] = cur
					ff[i][j] = true
				}
			}
		}
		t[i] = t[i-1][:st[top].x] + s[i][:st[top].y]
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &k)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
	}
	ok[n+1][0] = true
	for i := n; i > 0; i-- {
		for j := k; j >= 0; j-- {
			if ok[i][j] || ok[i+1][j] {
				ok[i][j] = true
			}
		}
		for j := k; j >= len(s[i]); j-- {
			if ok[i][j] || ok[i+1][j-len(s[i])] {
				ok[i][j] = true
			}
		}
	}
	solve()
	fmt.Println(t[n])
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
