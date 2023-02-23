package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007
const N = 1005

var n, cnt, res int
var sz [N]int
var v [][]int
var C [N][N]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	v = make([][]int, N)
	for i := 0; i < n-1; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		v[x] = append(v[x], y)
		v[y] = append(v[y], x)
	}
	for i := 0; i <= n; i++ {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j] + C[i-1][j-1]) % mod
		}
	}
	for i := 1; i <= n; i++ {
		cnt = 1
		dfs(i, 0)
		res += cnt
		res %= mod
	}
	fmt.Println(res * 500000004 % mod)
}

func dfs(x, f int) {
	sz[x] = 1
	for _, y := range v[x] {
		if y != f {
			dfs(y, x)
			sz[x] += sz[y]
			cnt *= C[sz[x]-1][sz[y]]
			cnt %= mod
		}
	}
}
