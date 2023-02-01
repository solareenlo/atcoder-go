package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF = 2100000000
const N = 200005

var m int
var s, t int
var f [N]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n, &m)
	sa := 0
	var b [N]int
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a, &b[i])
		b[i] -= a
		sa += a
	}
	for i := 0; i <= m; i++ {
		f[i] = INF
	}
	tmp := b[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	s, t = sa, sa
	cnt := 1
	f[sa] = 0
	b[n+1] = INF
	for i := 1; i <= n; i++ {
		if b[i] != b[i+1] {
			for t := 1; cnt > 0; t = min(t*2, cnt) {
				tic(b[i]*t, t)
				cnt -= t
			}
		}
		cnt++
	}
	for i := 0; i <= m; i++ {
		if f[i] == INF {
			fmt.Println(-1)
		} else {
			fmt.Println(f[i])
		}
	}
}

func tic(x, d int) {
	if x > 0 {
		t = min(m, t+x)
		for j := t; j >= s+x; j-- {
			f[j] = min(f[j], f[j-x]+d)
		}
	}
	if x < 0 {
		s = max(x+s, 0)
		for j := s; j <= t+x; j++ {
			f[j] = min(f[j], f[j-x]+d)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
