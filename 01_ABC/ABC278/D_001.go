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

	const N = 200005

	var n int
	fmt.Fscan(in, &n)

	m := make([]map[int]int, N)
	for i := range m {
		m[i] = make(map[int]int)
	}
	p := 0
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		m[p][i] = a
	}

	var q int
	fmt.Fscan(in, &q)
	c := 0
	for i := 1; i <= q; i++ {
		var o, x int
		fmt.Fscan(in, &o, &x)
		switch o {
		case 1:
			p++
			c = x
		case 2:
			var y int
			fmt.Fscan(in, &y)
			m[p][x] += y
		case 3:
			fmt.Fprintln(out, m[p][x]+c)
		}
	}
}
