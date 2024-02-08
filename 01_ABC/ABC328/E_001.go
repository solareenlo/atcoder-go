package main

import (
	"bufio"
	"fmt"
	"os"
)

var f [19]int
var u, v [109]int

func ff(xx int) int {
	if f[xx] == xx {
		return xx
	} else {
		return ff(f[xx])
	}
}

var n, m, k, ans int
var w [109]int

func dfs(x, y, num int) {
	if num == n-1 {
		ans = min(ans, y)
		return
	}
	if x > m {
		return
	}
	uu := ff(u[x])
	vv := ff(v[x])
	if uu != vv {
		f[uu] = vv
		dfs(x+1, (y+w[x])%k, num+1)
		f[uu] = uu
	}
	dfs(x+1, y, num)
	return
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m, &k)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &u[i], &v[i], &w[i])
	}
	for i := 1; i <= n; i++ {
		f[i] = i
	}
	ans = k
	dfs(1, 0, 0)
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
