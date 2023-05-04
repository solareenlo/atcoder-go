package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100
	const M = 299993

	v := 1
	for v*v%M != M-1 {
		v++
	}

	var n, z int
	fmt.Fscan(in, &n, &z)

	var v1, v2 [N + 5]int
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		y = v * y % M
		v1[i] = x - y
		v2[i] = x + y
	}

	var cnt [M + 5]int
	for i := 0; i < M; i++ {
		t := 1
		for j := 1; j <= n; j++ {
			t = t * (i - v1[j]) % M
		}
		if t < 0 {
			t += M
		}
		cnt[t]++
	}

	var inv [M + 5]int
	inv[1] = 1
	for i := 2; i < M; i++ {
		inv[i] = (M - M/i) * inv[M%i] % M
	}

	ans := 0
	for i := 0; i < M; i++ {
		t := 1
		for j := 1; j <= n; j++ {
			t = t * (i - v2[j]) % M
		}
		if t < 0 {
			t += M
		}
		t = inv[t]
		if t == 0 {
			if z == 0 {
				ans += M
			}
		} else {
			ans += cnt[z*t%M]
		}
	}
	fmt.Println(ans)
}
