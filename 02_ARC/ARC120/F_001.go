package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s, k, d int
	fmt.Fscan(in, &s, &k, &d)

	mod := 998244353
	inv := make([]int, 1000010)
	inv[1] = 1
	for i := 2; i <= s; i++ {
		inv[i] = (mod - mod/i) * inv[mod%i] % mod
	}
	t := 1
	for i := 1; i < k; i++ {
		t = t * inv[i] % mod * (s - k - i + 1) % mod
	}

	f := make([]int, s+1)
	for i := 0; i <= s && t != 0; i++ {
		f[i] = t
		t = t * inv[s-k-i] % mod * (k - 1 - i) % mod
	}
	for i := 1; (i << 1) <= s; i++ {
		tmp := f[i]
		if i&1 != 0 {
			tmp = mod - f[i]
		}
		f[i] = (tmp + f[i-1]) % mod
	}

	ans := 0
	for i := 1; i <= s; i++ {
		var a int
		fmt.Fscan(in, &a)
		ans = (ans + a*f[min(i-1, s-i)]) % mod
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
