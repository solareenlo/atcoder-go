package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

var (
	n int
	M = map[int]int{}
	u int
	d = [100005]int{}
)

func check(a, b int) int {
	if a == b {
		return 0
	}
	return 1
}

func as(a, b, c int) {
	t := a*n + b
	k := check(M[t], c)
	if M[t] == 0 {
		M[t] = c
		if M[t] != 0 {
			u++
		}
	} else if a == 2 {
		as(0, b+1, k+1)
	} else if a == 1 {
		d[b] = k + 1
	} else if k != 0 {
		fmt.Println(0)
		os.Exit(0)
	}
}

func exp(b, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 != 0 {
			res *= b
			res %= MOD
		}
		b *= b
		b %= MOD
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(in, &n, &m)

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		if a < b {
			a, b = b, a
		}
		as(a-b, b-1, c+1)
	}

	for i := 0; i < n; i++ {
		if d[i] != 0 && M[i] != 0 {
			as(0, i+1, check(M[i], d[i])+1)
			d[i] = 0
		}
	}

	for i := n - 2; i >= 0; i-- {
		if d[i] != 0 && M[i+1] != 0 {
			as(0, i, check(M[i+1], d[i])+1)
			d[i] = 0
		}
	}

	for i := 0; i < n; i++ {
		if d[i] != 0 {
			u++
		}
	}
	fmt.Println(exp(2, n*(n+1)/2-u))
}
