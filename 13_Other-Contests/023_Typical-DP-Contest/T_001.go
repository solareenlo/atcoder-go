package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100001
	const mod = 1000000007

	var k, n int
	fmt.Fscan(in, &k, &n)
	n--
	m := n
	x := make([]int, N)
	y := make([]int, N)
	x[1], y[0] = 1, 1

	var mul func([]int, []int, int)
	mul = func(x, y []int, n int) {
		tmp := make([]int, N)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				tmp[i+j] += x[i] * y[j] % mod
				tmp[i+j] %= mod
			}
		}
		for i := 2*n - 2; i >= n; i-- {
			for j := 1; j <= n; j++ {
				tmp[i-j] += tmp[i]
				tmp[i-j] %= mod
			}
		}
		for i := 0; i < n; i++ {
			x[i] = tmp[i]
		}
	}

	for m != 0 {
		if (m & 1) != 0 {
			mul(y, x, k)
		}
		mul(x, x, k)
		m >>= 1
	}
	ans := 0
	for i := 0; i < k; i++ {
		ans += y[i]
		ans %= mod
	}
	fmt.Println(ans)
}
