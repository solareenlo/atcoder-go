package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200005

	var a [N]int

	var n, d, p int
	fmt.Fscan(in, &n, &d, &p)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &f[i])
	}
	sort.Ints(f[1:])
	for i := n; i >= 1; i-- {
		a[i] = min(f[i]+a[i+1], p+a[min(d+i, n+1)])
	}
	fmt.Println(a[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
