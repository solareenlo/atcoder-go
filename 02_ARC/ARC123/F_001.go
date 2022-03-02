package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var (
	n  int
	mu = make([]int, 300003)
)

func get(a, b int) int {
	res := 0
	if a < b {
		a, b = b, a
	}
	r := 0
	for l := 1; l <= n; l = r + 1 {
		r = n / (n / l)
		val := 0
		d := n / l
		lim := d / a
		for i := 1; i <= lim; i++ {
			val += (d - i*a) / b
		}
		res += (mu[r] - mu[l-1]) * val
	}
	return res + 1
}

func print(a, b int) {
	defer out.Flush()
	if a+b > n {
		fmt.Fprint(out, b, " ")
		return
	}
	print(a, a+b)
	print(a+b, b)
}

func query(a, b, l, r, tot int) {
	if l == 1 && r == tot {
		print(a, b)
		return
	}
	v := get(a, a+b)
	if l <= v {
		query(a, a+b, l, min(r, v), v)
	}
	if r > v {
		query(a+b, b, max(1, l-v), r-v, tot-v)
	}
}

func main() {
	defer out.Flush()

	var a, b, l, r int
	fmt.Fscan(in, &a, &b, &n, &l, &r)

	inp := make([]int, n+1)
	pc := 0
	pri := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if inp[i] == 0 {
			pc++
			pri[pc] = i
			mu[i] = -1
		}
		for j := 1; j <= pc && i*pri[j] <= n; j++ {
			x := i * pri[j]
			inp[x] = 1
			if i%pri[j] == 0 {
				break
			}
			mu[x] = -mu[i]
		}
	}
	mu[1] = 1
	for i := 2; i <= n; i++ {
		mu[i] += mu[i-1]
	}
	if l == 1 {
		fmt.Fprint(out, a, " ")
		l++
	}
	l--
	r--
	query(a, b, l, r, get(a, b))
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
