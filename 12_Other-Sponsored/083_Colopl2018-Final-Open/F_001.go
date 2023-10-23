package main

import "fmt"

func calc(x, v, n int) int {
	for x < n && v >= (x+1) {
		x++
		v -= x
	}
	if x+v*2 <= n {
		x += v * 2
		v = 0
		for 3*(x+1) <= n {
			x = 3 * (x + 1)
		}
		if x == n {
			return 0
		}
		v = x + 1
		x++
	}
	d := (n - x) / 2
	v -= d
	x += d * 2
	for x < n {
		x++
		if v >= x {
			v -= x
		} else {
			v += x
		}
	}
	return v
}

func main() {
	var n, l, r int
	fmt.Scan(&n, &l, &r)
	ed := calc(0, 0, n)
	ans := 0
	for x := 0; x < n; x = 3 * (x + 1) {
		R := min(x-2, (n-3-x)/2)
		if R >= 1 {
			tmp := calc(x+3+R*2, x-2-R, n)
			if tmp >= l && tmp <= r {
				ans = ans + R
			}
		}
		R++
		if R < 1 {
			R = 1
		}
		for t, k := x+1+R*2, R; t <= n && t <= 3*(x+1); t, k = t+2, k+1 {
			tmp := calc(t, 3*(x+1+k), n)
			if tmp >= l && tmp <= r {
				ans++
			}
		}
		if ed >= l && ed <= r {
			u := min(x+2, (n-x)/2+1)
			ans = ans + u
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
