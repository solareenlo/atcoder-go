package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m int
var a [15][15]int
var cnt map[int]int
var ans int

func dfs(i, j int) {
	if i < 1 || i > n || j < 1 || j > m || cnt[a[i][j]] > 0 {
		return
	}
	cnt[a[i][j]]++
	if i == n && j == m {
		ans++
	}
	dfs(i+1, j)
	dfs(i, j+1)
	cnt[a[i][j]]--
}

func main() {
	in := bufio.NewReader(os.Stdin)

	cnt = make(map[int]int)

	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	dfs(1, 1)
	fmt.Println(ans)
}
