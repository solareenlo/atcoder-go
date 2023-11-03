package main

import (
	"bufio"
	"fmt"
	"os"
)

var f [200200]int

func find(x int) int {
	if f[x] == x {
		return x
	}
	f[x] = find(f[x])
	return f[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var p [200200]int

	var n int
	fmt.Fscan(in, &n)
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}
	for i := 1; i <= n; i++ {
		f[i] = i
	}
	var q int
	fmt.Fscan(in, &q)
	for i := 1; i <= q; i++ {
		var j, k int
		fmt.Fscan(in, &j, &k)
		if j == 1 {
			fmt.Fscan(in, &j)
			for find(j) != find(k) {
				f[find(k)] = find(p[find(k)])
			}
		} else {
			fmt.Fprintln(out, find(k))
		}
	}
}
