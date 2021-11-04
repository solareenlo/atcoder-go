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

	a := make([]int, n+1)
	f := make([]int, n+1)
	sum := 0
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i])
		if i&1 != 0 {
			sum += a[i]
		}
		if i == 1 {
			continue
		}
		if i&1 != 0 {
			f[i] = max(f[i-2]+a[i], f[i-1])
		} else {
			f[i] = max(f[i-2]+a[i], sum)
		}
	}

	fmt.Println(f[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
