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

	const MX = 500005

	var n, q int
	fmt.Fscan(in, &n, &q)

	var e [MX]int
	for i := 0; i < n; i++ {
		var d int
		fmt.Fscan(in, &d)
		e[i+1] = e[i] + d
	}

	for i := 0; i < q; i++ {
		var t, l, r int
		fmt.Fscan(in, &t, &l, &r)
		fmt.Fprintln(out, upperBound(e[:n+1], t-l)-lowerBound(e[:n+1], t-r))
	}
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
