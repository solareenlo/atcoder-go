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

	const MOD = 1000000007

	var n int
	fmt.Fscan(in, &n)
	var x, y [2000]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	var p2 [2001]int
	p2[0] = 1
	for i := 1; i <= n; i++ {
		p2[i] = p2[i-1] * 2 % MOD
	}
	ans := 0
	ind := make([]int, n)
	var t [2000]float64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < i {
				ind[j] = j
				t[j] = math.Atan2(float64(y[j]-y[i]), float64(x[j]-x[i]))
			} else if j > i {
				ind[j-1] = j
				t[j] = math.Atan2(float64(y[j]-y[i]), float64(x[j]-x[i]))
			}
		}
		tmp := ind[:n-1]
		sort.Slice(tmp, func(p, q int) bool {
			return t[tmp[p]] < t[tmp[q]]
		})
		r := 1
		for l := 0; l < n-1; l++ {
			p := ind[l]
			for r != l {
				q := ind[r]
				if (x[p]-x[i])*(y[q]-y[i])-(x[q]-x[i])*(y[p]-y[i]) < 0 {
					break
				}
				r = (r + 1) % (n - 1)
			}
			c := r - l - 1
			if c < 0 {
				c += (n - 1)
			}
			ans += (p2[c] - 1) * (x[i]*y[p]%MOD - x[p]*y[i]%MOD + 2*MOD)
			ans %= MOD
		}
	}
	fmt.Println(ans)
}
