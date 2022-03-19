package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	c := make([]int, (n+1)*(n+1))
	type pair struct{ a, b int }
	d := make([]pair, n+1)
	d[0].a = -1
	d[0].b = -1
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &d[i].a)
		a[i] = d[i].a
		c[d[i].a] = i
		d[i].b = i
	}
	sort.Slice(d, func(i, j int) bool {
		return d[i].a < d[j].a
	})

	for i, p := 1, 1; i <= n; i++ {
		for j := 1; j < d[i].b; j++ {
			for c[p] != 0 {
				p++
			}
			c[p] = d[i].b
		}
		if p > a[d[i].b] {
			fmt.Fprintln(out, "No")
			return
		}
	}

	for i, p := n, n*n; i >= 1; i-- {
		for j := d[i].b + 1; j <= n; j++ {
			for c[p] != 0 {
				p--
			}
			c[p] = d[i].b
		}
		if p < a[d[i].b] {
			fmt.Fprintln(out, "No")
			return
		}
	}

	fmt.Fprintln(out, "Yes")
	for i := 1; i <= n*n; i++ {
		fmt.Fprint(out, c[i], " ")
	}
}
