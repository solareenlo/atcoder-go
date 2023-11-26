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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n+1)
	b := make([]int, m+1)
	c := make([]int, n+1)
	d := make([]int, m+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &b[i])
	}
	i, j, t := 1, 1, 0
	for i <= n || j <= m {
		if i == n+1 {
			t++
			d[j] = t
			j++
		} else if j == m+1 {
			t++
			c[i] = t
			i++
		} else if a[i] < b[j] {
			t++
			c[i] = t
			i++
		} else {
			t++
			d[j] = t
			j++
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d ", c[i])
	}
	fmt.Fprintln(out)
	for i := 1; i <= m; i++ {
		fmt.Fprintf(out, "%d ", d[i])
	}
	fmt.Fprintln(out)
}
