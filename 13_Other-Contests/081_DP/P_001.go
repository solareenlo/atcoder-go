package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007

var g [100010][]int
var f [100010][3]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	DFs(1, 0)
	fmt.Println((f[1][0] + f[1][1]) % mod)
}

func DFs(u, fa int) {
	f[u][0] = 1
	f[u][1] = 1
	for _, v := range g[u] {
		if v == fa {
			continue
		}
		DFs(v, u)
		f[u][0] = f[u][0] * f[v][1] % mod
		f[u][1] = f[u][1] * (f[v][1] + f[v][0]) % mod
	}
	return
}
