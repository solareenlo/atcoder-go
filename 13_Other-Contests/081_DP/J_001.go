package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var f [311][311][311]float64

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	var c [5]int
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		c[x]++
	}
	fmt.Println(Dfs(c[1], c[2], c[3]))
}

func Dfs(x, y, z int) float64 {
	if x == 0 && y == 0 && z == 0 {
		return 0.0
	}
	if f[x][y][z] != 0.0 {
		return f[x][y][z]
	}
	f[x][y][z] += float64(n) / float64(x+y+z)
	if z != 0 {
		f[x][y][z] += Dfs(x, y+1, z-1) * float64(z) / float64(x+y+z)
	}
	if y != 0 {
		f[x][y][z] += Dfs(x+1, y-1, z) * float64(y) / float64(x+y+z)
	}
	if x != 0 {
		f[x][y][z] += Dfs(x-1, y, z) * float64(x) / float64(x+y+z)
	}
	return f[x][y][z]
}
