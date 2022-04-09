package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100005

var (
	used = make([]bool, N)
	ans  = make([]int, 0)
	s    = make([]int, N)
	G    = make([][]int, N)
)

func dfs(v int) {
	used[v] = true
	ans = append(ans, v)
	s[v] ^= 1
	for _, u := range G[v] {
		if used[u] {
			continue
		}
		dfs(u)
		ans = append(ans, v)
		s[v] ^= 1
		if s[u] == 1 {
			ans = append(ans, u)
			s[u] ^= 1
			ans = append(ans, v)
			s[v] ^= 1
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}

	var S string
	fmt.Fscan(in, &S)
	for i := 1; i <= n; i++ {
		s[i] = int(S[i-1] - '0')
	}

	dfs(1)
	if s[1] == 1 {
		ans = append(ans, G[1][0])
		ans = append(ans, 1)
		ans = append(ans, G[1][0])
	}

	fmt.Fprintln(out, len(ans))
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
