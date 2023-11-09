package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 300005
	const mod = 998244353

	var a, b [N]int

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	m--
	if m+k > n {
		for i := 0; i < n; i++ {
			b[i] = a[(k+m+i)%n]
		}
		for i := 0; i < n-m; i++ {
			a[(i+k)%(n-m)] = b[i]
		}
		for i := n - m; i < n; i++ {
			a[i] = b[i]
		}
		k = n - m
	}
	ans := 1
	for i := 1; i <= m; i++ {
		ans = ans * i % mod
	}
	mx := 0
	for i := 0; i < k; i++ {
		if a[i] > mx {
			ans = ans * (m + 1) % mod
			mx = a[i]
		}
	}
	for i := k; i < k+m; i++ {
		if a[i] < mx {
			ans = 0
		} else {
			mx = a[i]
		}
	}
	fmt.Println(ans)
}
