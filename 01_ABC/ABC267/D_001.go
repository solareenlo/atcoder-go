package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200005

	var n, m int
	fmt.Fscan(in, &n, &m)

	var a [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var f [N]int
	for i := range f {
		f[i] = -128_000_000_000_000_000
	}
	f[0] = 0
	for i := 1; i <= n; i++ {
		for j := m; j > 0; j-- {
			f[j] = max(f[j], f[j-1]+a[i]*j)
		}
	}
	fmt.Println(f[m])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
