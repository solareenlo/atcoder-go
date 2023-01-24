package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type node struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	const N = 1010
	c := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i])
	}

	d := make([]node, N)
	var x int
	for i := 1; i <= n; i++ {
		s, q := 0, 0
		for j := 6; j > 0; j-- {
			fmt.Fscan(in, &x)
			s += x
			q += x * x
		}
		d[i].y = 6*q - s*s - 36*c[i]
		d[i].x = s
	}
	tmp := d[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].y == tmp[j].y {
			return tmp[i].x < tmp[j].x
		}
		return tmp[i].y < tmp[j].y
	})

	s := make([]int, N)
	sx, sy, res := 0, 0, 0
	for i := n - k + 1; i <= n; i++ {
		s[i] = 1
		sx += d[i].x
		sy += d[i].y
	}

	e := make([]node, 0)
	res = sx*sx + sy
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if d[i].x > d[j].x {
				e = append(e, node{i, j})
			}
		}
	}
	sort.Slice(e, func(i, j int) bool {
		return (d[e[i].x].y-d[e[i].y].y)*(d[e[j].y].x-d[e[j].x].x) < (d[e[j].x].y-d[e[j].y].y)*(d[e[i].y].x-d[e[i].x].x)
	})

	for i, _ := range e {
		if s[e[i].x] == 0 && s[e[i].y] != 0 {
			s[e[i].x]++
			s[e[i].y]--
			sx += d[e[i].x].x - d[e[i].y].x
			sy += d[e[i].x].y - d[e[i].y].y
			res = max(res, sx*sx+sy)
		}
	}

	fmt.Println((res%MOD + MOD) * powMod(36, MOD-2) % MOD)
}

const MOD = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
