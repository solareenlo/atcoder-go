package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 2100

type ask struct{ x, y, z int }

var (
	f = [N]int{}
	a = make([]ask, N)
)

func find(x int) int {
	if f[x] != 0 {
		f[x] = find(f[x])
		return f[x]
	}
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, x int
	fmt.Fscan(in, &n, &m, &x)

	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &a[i].x, &a[i].y, &a[i].z)
	}

	tmp := a[1 : m+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].z < tmp[j].z
	})

	mini := x
	ans := 1
	cnt := 0
	mod := 1_000_000_007
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			f[j] = 0
		}
		r := a[i].z
		f[a[i].x] = a[i].y
		for j := 1; j <= m; j++ {
			x := find(a[j].x)
			y := find(a[j].y)
			if x == y {
				continue
			}
			f[x] = y
			r += a[j].z
		}
		mini = min(mini, r)
		if r > x {
			ans = ans * 2 % mod
		}
		if r == x {
			cnt++
		}
	}

	if mini == x {
		cnt--
	}

	tt := 1
	for i := 1; i <= cnt; i++ {
		tt = tt * 2 % mod
	}
	tt = (tt - 1 + mod) % mod

	if cnt == 0 && mini != x {
		fmt.Println(0)
	} else {
		fmt.Println(ans * tt % mod * 2 % mod)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
