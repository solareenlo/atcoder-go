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

	var sum [N]int

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		if i%2 == 1 {
			sum[i] = sum[i-1] + 1*(a[i]-a[i-1])
		} else {
			sum[i] = sum[i-1] + 0*(a[i]-a[i-1])
		}
	}
	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var l, r int
		fmt.Fscan(in, &l, &r)
		x := upperBound(a[1:], l) + 1
		y := lowerBound(a[1:], r) + 1
		tmp0 := 0
		tmp1 := 0
		if x%2 == 1 {
			tmp0 = 1
		}
		if y%2 == 1 {
			tmp1 = 1
		}
		fmt.Fprintln(out, sum[y]-sum[x]+tmp0*(a[x]-l)-tmp1*(a[y]-r))
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
