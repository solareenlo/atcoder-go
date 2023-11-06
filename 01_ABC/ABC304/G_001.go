package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m int
var sz [10100000]int
var ch [10100100][2]int

func gz(x, y, d int) int {
	if x == 0 || y == 0 {
		return 0
	}
	if d == -1 {
		return min(sz[x], sz[y])
	}
	if ((m >> d) & 1) != 0 {
		return gz(ch[x][0], ch[y][1], d-1) + gz(ch[x][1], ch[y][0], d-1)
	}
	m1 := min(sz[ch[x][1]], sz[ch[y][0]])
	m2 := min(sz[ch[x][0]], sz[ch[y][1]])
	mp := m1 + m2
	mp += min(gz(ch[x][0], ch[y][0], d-1), min(sz[ch[x][0]]-m2, sz[ch[y][0]]-m1))
	mp += min(gz(ch[x][1], ch[y][1], d-1), min(sz[ch[x][1]]-m1, sz[ch[y][1]]-m2))
	return mp
}

func sol(x, d int) int {
	if x == 0 {
		return 0
	}
	if d == -1 {
		return sz[x] / 2
	}
	if ((m >> d) & 1) == 0 {
		a := ch[x][0]
		b := ch[x][1]
		if sz[a] < sz[b] {
			a, b = b, a
		}
		return sz[b] + min(sz[a]-sz[b], sol(a, d-1))
	}
	return gz(ch[x][0], ch[x][1], d-1)
}

func ck() bool { return sol(1, 29) >= (n+1)/2 }

func main() {
	in := bufio.NewReader(os.Stdin)

	var a [201000]int

	fmt.Fscan(in, &n)
	cn := 1
	for i := 1; i <= n*2; i++ {
		fmt.Fscan(in, &a[i])
		u := 1
		sz[u]++
		for j := 30 - 1; j >= 0; j-- {
			p := (a[i] >> j) & 1
			if ch[u][p] == 0 {
				cn++
				ch[u][p] = cn
			}
			u = ch[u][p]
			sz[u]++
		}
	}
	l := 0
	r := ((1 << 30) - 1)
	ans := 0
	for l <= r {
		m = (l + r) >> 1
		if ck() {
			ans = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
