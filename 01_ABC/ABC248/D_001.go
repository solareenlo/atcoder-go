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

	v := make([][]int, 210000)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		v[x] = append(v[x], i)
	}

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var l, r, x int
		fmt.Fscan(in, &l, &r, &x)
		fmt.Fprintln(out, upperBound(v[x], r)-lowerBound(v[x], l))
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
