package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200200

var dp [N]int
var e [N][]int

func dfs(x int) int {
	if dp[x] != 0 {
		return dp[x]
	}
	for _, i := range e[x] {
		dp[x] = max(dp[x], dfs(i))
	}
	dp[x]++
	return dp[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		e[u] = append(e[u], v)
	}
	k := 0
	for i := 1; i <= n; i++ {
		k = max(k, dfs(i))
	}
	if k < n {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
	for i := 1; i <= n; i++ {
		fmt.Println(n + 1 - dp[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
