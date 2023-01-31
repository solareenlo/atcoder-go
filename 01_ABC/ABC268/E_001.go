package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	var a [400040]int
	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(in, &t)
		t = (t + n - i) % n
		a[t]++
		a[t+n]++
		a[t+n/2]--
		a[t+(n+1)/2]--
	}
	for i := 1; i <= 2*n; i++ {
		a[i] = a[i-1] + a[i]
	}
	for i := 1; i <= 2*n; i++ {
		a[i] = a[i-1] + a[i]
	}
	mn := 1_000_000_000_000_000
	for i := 0; i < n; i++ {
		mn = min(mn, a[i]+a[i+n])
	}
	fmt.Println(mn)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
