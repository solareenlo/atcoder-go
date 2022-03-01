package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, 10005)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	if n&1 != 0 {
		n++
		a[n] = 0
	}
	tmp := a[1 : n+1]
	sort.Ints(tmp)

	f1 := -1 << 60
	f2 := 1 << 60
	for j := 1; j <= (n >> 1); j++ {
		X := a[j] + a[n+1-j]
		f1 = max(f1, X)
		f2 = min(f2, X)
	}

	ans := f1 - f2
	for i := 2; i <= n; i += 2 {
		j := n + i - 2
		for a[j] > 0 {
			a[j+2] = a[j]
			j--
		}
		a[j+1] = 0
		a[j+2] = 0
		f1 = -2e9
		f2 = 2e9
		for j := 1; j <= ((n + i) >> 1); j++ {
			X := a[j] + a[n+i+1-j]
			f1 = max(f1, X)
			f2 = min(f2, X)
		}
		ans = min(ans, f1-f2)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
