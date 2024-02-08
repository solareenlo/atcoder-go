package main

import (
	"bufio"
	"fmt"
	"os"
)

var f, dis [1000100]int

func find(x int) int {
	if f[x] != x {
		res := f[x]
		f[x] = find(f[x])
		dis[x] += dis[res]
	}
	return f[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)
	for i := 1; i <= n; i++ {
		f[i] = i
	}
	for i := 1; i <= q; i++ {
		var x, y, z int
		fmt.Fscan(in, &x, &y, &z)
		r1 := find(x)
		r2 := find(y)
		if !(r1 == r2 && dis[x]-dis[y] != z) {
			fmt.Fprintf(out, "%d ", i)
			f[r1] = r2
			dis[r1] = z + dis[y] - dis[x]
		}
	}
}
