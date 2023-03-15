package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var L, n, k int
	fmt.Fscan(in, &L, &n, &k)

	x := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i])
	}

	a := make([]float64, n)
	for i := 0; i < n-1; i++ {
		a[i] = x[i+1] - x[i]
	}
	a[n-1] = float64(L) - x[n-1]

	sort.Float64s(a)

	lo := 0.0
	hi := 1e9

	for (hi - lo) > math.Max(1., lo)*1e-8 {
		x := (lo + hi) / 2
		cnt1 := 0
		cnt2 := k + 1
		for i := 0; i < n; i++ {
			cost := int(math.Min(1.0*float64(k)+1.0, math.Ceil(1.0*a[i]/x)-1))
			cnt1 += cost
		}
		i := 0
		for i < n && a[i] <= x {
			i++
		}
		if i < n {
			se := 2.0*x - a[i]
			if se > 0 {
				cnt2 = 0
				for j := 0; j < n; j++ {
					if j == i {
						continue
					}
					cost := int(math.Min(1.0*float64(k)+1.0, math.Ceil(1.*a[j]/se)-1))
					cnt2 += cost
				}
			}
		}
		if min(cnt1, cnt2) <= k {
			hi = x
		} else {
			lo = x
		}
	}

	fmt.Println(lo)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
