package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)
	const MX = 100005

	var n, m int
	var s, t string
	fmt.Fscan(in, &n, &m, &s, &t)

	var num func(int) int
	num = func(p int) int {
		x := int(s[p] - '0')
		if x < 0 || x > 9 {
			return -1
		}
		return x
	}

	c := make([][]int, MX)
	for i := range c {
		c[i] = make([]int, 2)
	}
	cur := 0
	for i := 0; i < m; i++ {
		for j := 0; j < 2; j++ {
			x := 0
			h := 1
			for num(cur) == -1 {
				if s[cur] == '~' {
					h = -1
				}
				cur++
			}
			for num(cur) != -1 {
				x = x*10 + num(cur)
				cur++
			}
			c[i][j] = x * h
		}
	}
	f := make([]bool, MX)
	for i := 0; i < n; i++ {
		f[i+1] = (t[i] == 'T')
	}

	var ok func(int) bool
	ok = func(i int) bool {
		h := (i < 0)
		x := i
		if h {
			x = -i
		}
		return (f[x] && !h) || (!f[x] && h)
	}
	var flip func(int)
	flip = func(x int) {
		x = abs(x)
		f[x] = !f[x]
	}
	ans := INF
	var dfs func(int)
	dfs = func(dep int) {
		w := -1
		for i := 0; i < m; i++ {
			if !ok(c[i][0]) && !ok(c[i][1]) {
				w = i
				break
			}
		}
		if w == -1 {
			ans = min(ans, dep)
			return
		}
		if dep == 10 {
			return
		}
		dep++
		flip(c[w][0])
		dfs(dep)
		flip(c[w][0])
		flip(c[w][1])
		dfs(dep)
		flip(c[w][1])
	}
	dfs(0)

	if ans == INF {
		fmt.Println("TOO LARGE")
	} else {
		fmt.Println(ans)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
