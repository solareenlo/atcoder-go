package main

import (
	"bufio"
	"fmt"
	"os"
)

const P = 998244353

var n, m, a int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m, &a)

	val := make([]int, 300005)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &val[i])
	}
	gg := gcd(a, m)
	a /= gg
	m /= gg
	sum := 0
	for i := 1; i <= n; i++ {
		val[i] /= gg
		sum ^= sg(val[i])
	}
	if sum == 0 {
		fmt.Println(0)
		return
	}
	ans := 0
	for i := 1; i <= n; i++ {
		ans = (ans + get(val[i], sum^sg(val[i]))) % P
	}
	fmt.Println(ans)
}

func sg(x int) int {
	if a != 0 {
		return min(calc(x), x/a)
	}
	return x
}

func get(x, v int) int {
	if a == 0 {
		if x > v {
			return 1
		}
		return 0
	}
	if x < a {
		return 0
	}
	y := x - a
	l := (y % m / a) & 1
	r := calc(y)
	vv := min(y/a, r)
	if v < l || v > vv {
		return 0
	}
	if v < vv {
		return 1
	}
	if vv >= r {
		return (y / m) - (r - l) + 1
	}
	return 1
}

func calc(x int) int {
	if x < a {
		return 0
	}
	if a*2 <= m {
		if x < m {
			return (x / a) & 1
		}
		x %= m
		tt := (m/a+1)*a - m
		l := 0
		if ((m/a + 1) & 1) != 0 {
			l = tt
			tt = a - tt
		}
		if x >= l+a && x <= l+a+tt-1 {
			return 2
		}
		if x < l+a {
			x += m
		}
		return (x / a) & 1
	}
	return (m - (x-a)%m - 1) / (m - a)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
