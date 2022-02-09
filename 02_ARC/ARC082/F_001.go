package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var x, n int
	fmt.Fscan(in, &x, &n)

	r := make([]int, 100100)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &r[i])
	}

	var m int
	fmt.Fscan(in, &m)

	u := x
	tag := -1
	p, s, d := 0, 0, 0
	for i := 0; i < m; i++ {
		var y, z int
		fmt.Fscan(in, &y, &z)
		for p < n && r[p+1] <= y {
			v := tag * (r[p+1] - r[p])
			d = min(x, max(0, d+v))
			u = min(x, max(0, u+v))
			s += v
			tag = -tag
			p++
		}
		v := tag * (y - r[p])
		fmt.Fprintln(out, min(min(x, max(0, u+v)), max(min(x, max(0, d+v)), s+v+z)))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
