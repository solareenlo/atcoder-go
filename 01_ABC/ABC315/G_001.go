package main

import "fmt"

func main() {
	var n, a, b, c, m int
	fmt.Scan(&n, &a, &b, &c, &m)
	r := gcd(a, b)
	a /= r
	b /= r
	res := 0
	_, a2, _ := gcdExt(a, b)
	for i := 1; i <= n && a+b+i*c <= m; i++ {
		x := m - i*c
		if x%r != 0 {
			continue
		}
		x /= r
		maxx := min(n, (x-b)/a)
		minn := max(0, (x-n*b-1)/a)
		if minn > maxx {
			continue
		}
		x = x % b * a2 % b
		var tmp0 int
		if maxx < x {
			tmp0 = 0
		} else {
			tmp0 = (maxx-x)/b + 1
		}
		var tmp1 int
		if minn < x {
			tmp1 = 0
		} else {
			tmp1 = (minn-x)/b + 1
		}
		res += tmp0 - tmp1
	}
	fmt.Println(res)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func gcdExt(x, y int) (int, int, int) {
	if y == 0 {
		return x, 1, 0
	}
	g, a, b := gcdExt(y, x%y)
	return g, b, a - x/y*b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
