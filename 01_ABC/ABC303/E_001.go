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

	const N = 200200

	d := make([]int, N)
	a := make([]int, N)

	var n int
	fmt.Fscan(in, &n)
	s := n
	for i := 1; i < n; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		d[u]++
		d[v]++
	}
	c := 0
	for i := 1; i <= n; i++ {
		if d[i] >= 3 {
			c++
			a[c] = d[i]
			s -= d[i] + 1
		}
	}
	for i := 1; i <= s/3; i++ {
		c++
		a[c] = 2
	}
	tmp := a[1 : c+1]
	sort.Ints(tmp)
	for i := 1; i <= c; i++ {
		fmt.Fprintf(out, "%d ", a[i])
	}
}
