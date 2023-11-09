package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 110

var n, tim int
var e [N][N]int
var Vis, match [N]int

func Dfs(u int) int {
	for i := 1; i <= n; i++ {
		if e[u][i] != 0 && Vis[i] != tim {
			Vis[i] = tim
			if match[i] == 0 || Dfs(match[i]) != 0 {
				match[i] = u
				return 1
			}
		}
	}
	return 0
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var ans [N][N]int

	var m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			var x int
			fmt.Fscan(in, &x)
			e[i][x]++
		}
	}
	for i := 1; i <= m; i++ {
		res := 0
		for j := 1; j <= n; j++ {
			tim++
			if Dfs(j) == 0 {
				fmt.Println("No")
				return
			}
			res++
		}
		for j := 1; j <= n; j++ {
			ans[match[j]][i] = j
			e[match[j]][j]--
			match[j] = 0
		}
	}
	fmt.Println("Yes")
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Printf("%d ", ans[i][j])
		}
		fmt.Println()
	}
}
