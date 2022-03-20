package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	g = make([][]int, 100005)
	a = make([]int, 100005)
)

func dfs(x, y int) int {
	if len(g[x]) < 2 {
		return a[x]
	}
	s := 0
	mx := 0
	for _, i := range g[x] {
		if i != y {
			u := dfs(i, x)
			mx = max(mx, u)
			s += u
		}
	}
	w := s - a[x]
	f := s - 2*w
	if w < 0 || f < 0 || w > s-mx {
		fmt.Println("NO")
		os.Exit(0)
	}
	return f
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	if n == 2 {
		if a[1] == a[2] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		os.Exit(0)
	}

	for i := 1; i < n; i++ {
		var j, k int
		fmt.Fscan(in, &j, &k)
		g[j] = append(g[j], k)
		g[k] = append(g[k], j)
	}

	for i := 1; ; i++ {
		if len(g[i]) > 1 {
			if dfs(i, 0) != 0 {
				fmt.Println("NO")
			} else {
				fmt.Println("YES")
			}
			os.Exit(0)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
