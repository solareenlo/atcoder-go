package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 114514

var (
	m   int
	ans int
	v   = make([][]int, N)
)

func dfs(x, f int) int {
	p := 0
	for _, y := range v[x] {
		if y != f {
			p = max(p, dfs(y, x))
		}
	}
	p++
	if p >= m && f != 1 {
		ans++
	}
	return p % m
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, d int
	fmt.Fscan(in, &n, &m, &d)
	if d != 1 {
		ans++
	}

	for i := 2; i < n+1; i++ {
		var d int
		fmt.Fscan(in, &d)
		v[d] = append(v[d], i)
	}
	dfs(1, 1)
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
