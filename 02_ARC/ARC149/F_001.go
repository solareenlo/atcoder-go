package main

import (
	"fmt"
	"os"
)

var p, q, n, l, r int

func fir_son(x int) int {
	if x == 0 {
		return 1
	}
	t := (x + q - 1) / q
	if t <= n/p {
		return t * p
	}
	return n + 1
}

func calc(l, r int) int {
	ans := 0
	for l < r {
		ans += r - l
		l = fir_son(l)
		r = fir_son(r)
	}
	return ans
}

func dfs(x int) {
	if r <= 0 {
		os.Exit(0)
	}
	l--
	r--
	if l <= 0 {
		fmt.Println(x)
	}
	if x%q != 0 {
		return
	}
	fir := fir_son(x)
	var lst int
	if x == 0 {
		lst = min(p, n+1) - 1
	} else {
		lst = min(fir+p, n+1) - 1
	}
	p := fir
	if l > 0 {
		lb := fir + 1
		rb := lst + 1
		for lb < rb {
			mid := (lb + rb) / 2
			if calc(fir, mid) < l {
				lb = mid + 1
				p = mid
			} else {
				rb = mid
			}
		}
	}
	if p > fir {
		tmp := calc(fir, p)
		l -= tmp
		r -= tmp
	}
	for i := p; i <= lst; i++ {
		dfs(i)
	}
}

func main() {
	fmt.Scan(&p, &q)
	fmt.Scan(&n, &l, &r)
	l++
	r++
	dfs(0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
