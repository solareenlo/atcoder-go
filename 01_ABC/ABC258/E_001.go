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

	var n, q, x int
	fmt.Fscan(in, &n, &q, &x)

	const N = 200010
	w := make([]int, N*2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &w[i])
		w[i+n] = w[i]
	}

	for i := 1; i <= 2*n; i++ {
		w[i] += w[i-1]
	}
	v := make([]int, N)
	a := make([]int, N)
	t, cnt := 0, 1
	for i := 1; v[t] == 0; i++ {
		v[t] = i
		u := lowerBound(w[1:2*n+1], w[t]+x%w[n]) + 1
		a[i] = u - t + x/w[n]*n
		t = u % n
		cnt++
	}

	o := cnt - v[t]
	p := v[t]
	var k int
	for q > 0 {
		fmt.Fscan(in, &k)
		if k > p {
			k = (k-p)%o + p
		}
		fmt.Fprintln(out, a[k])
		q--
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
