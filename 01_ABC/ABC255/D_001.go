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

	const N = 200010
	a := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	tmp := a[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	sum := make([]int, N)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + a[i]
	}

	for q > 0 {
		var x int
		fmt.Fscan(in, &x)
		tmp := a[1 : n+1]
		k := upperBound(tmp, x)
		fmt.Fprintln(out, (x*k-sum[k])+((sum[n]-sum[k])-x*(n-k)))
		q--
	}
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
