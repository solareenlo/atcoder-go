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

	var n, q int
	fmt.Fscan(in, &n, &q)
	var ac [100005]int
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i])
		ac[i+1] = ac[i] + x[i]
	}
	for i := 0; i < q; i++ {
		var c, d int
		fmt.Fscan(in, &c, &d)
		l := lowerBound(x, c-d)
		m := lowerBound(x, c)
		r := lowerBound(x, c+d)
		fmt.Fprintln(out, c*(2*m-l-r)-(2*ac[m]-ac[l]-ac[r])+d*(n-r+l))
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
