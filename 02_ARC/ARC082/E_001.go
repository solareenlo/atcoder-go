package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mod := 998244353

	p := make([]int, 210)
	x := make([]int, 210)
	y := make([]int, 210)
	p[0] = 1
	for i := 1; i <= n; i++ {
		fmt.Scan(&x[i], &y[i])
		p[i] = p[i-1] * 2 % mod
	}

	ans := p[n] - n - 1
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			d := 0
			for k := j + 1; k <= n; k++ {
				if (y[i]-y[j])*(x[i]-x[k]) == (y[i]-y[k])*(x[i]-x[j]) {
					d++
				}
			}
			ans = (ans + mod - p[d]) % mod
		}
	}
	fmt.Println(ans)
}
