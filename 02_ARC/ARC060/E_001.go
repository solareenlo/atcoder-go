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

	x := make([]int, n)
	for i := range x {
		fmt.Fscan(in, &x[i])
	}

	var L int
	fmt.Fscan(in, &L)
	r := [18][100000]int{}
	for i := 0; i < n; i++ {
		r[0][i] = upperBound(x, x[i]+L) - 1
	}
	for i := 0; i < 17; i++ {
		for j := 0; j < n; j++ {
			r[i+1][j] = r[i][r[i][j]]
		}
	}

	var Q int
	fmt.Fscan(in, &Q)
	for j := 0; j < Q; j++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		if a > b {
			a, b = b, a
		}
		res := 0
		N := a
		for i := 17; i >= 0; i-- {
			if r[i][N] < b {
				N = r[i][N]
				res += 1 << i
			}
		}
		if N < b {
			res++
		}
		fmt.Fprintln(out, res)
	}
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
