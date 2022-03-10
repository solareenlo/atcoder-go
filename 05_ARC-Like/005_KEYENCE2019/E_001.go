package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, d int
	fmt.Fscan(in, &n, &d)

	a := make([]int, n+1)
	b := make([]int, n+1)
	mini := 1 << 60
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		b[i] = a[i]
		mini = min(mini, a[i])
	}

	for i := 2; i <= n; i++ {
		b[i] = min(b[i], b[i-1]+d)
	}

	for i := n - 1; i > 0; i-- {
		b[i] = min(b[i], b[i+1]+d)
	}

	p := 0
	for i := 0; i < n+1; i++ {
		if mini == a[i] {
			p = i
			break
		}
	}

	A := 0
	for i := p - 1; i > 0; i-- {
		A += a[i] + b[i+1] + d
	}

	for i := p; i < n; i++ {
		A += b[i] + a[i+1] + d
	}
	fmt.Println(A)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
