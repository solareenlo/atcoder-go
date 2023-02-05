package main

import (
	"bufio"
	"fmt"
	"os"
)

var v map[int][]int
var vis map[int]bool
var ans int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	v = make(map[int][]int)
	vis = make(map[int]bool)

	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if _, ok := v[x]; !ok {
			v[x] = make([]int, 0)
		}
		v[x] = append(v[x], y)
		if _, ok := v[y]; !ok {
			v[y] = make([]int, 0)
		}
		v[y] = append(v[y], x)
	}
	dfs(1)
	fmt.Println(ans)
}

func dfs(x int) {
	vis[x] = true
	ans = max(ans, x)
	for _, y := range v[x] {
		if !vis[y] {
			dfs(y)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
