package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 998244353

	var inv [200005]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a[1:])
	s := 1
	ans := 1
	for i, j := 1, n; i <= j; {
		ans = ans * s % mod
		if a[i]+a[j] < m {
			i++
			s--
		} else {
			j--
			s++
		}
	}
	inv[0] = 1
	inv[1] = 1
	for i := 2; i <= n; i++ {
		inv[i] = (mod - mod/i) * inv[mod%i] % mod
	}
	for i := 2; i <= n; i++ {
		inv[i] = inv[i-1] * inv[i] % mod
	}
	for i, j := 1, 0; i <= n; i++ {
		if (i == n) || (a[i] != a[i+1]) {
			ans = ans * inv[i-j] % mod
			j = i
		}
	}
	fmt.Println(ans)
}
