package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		fmt.Fscan(in, &a[i])
	}
	c := make([]int, n+m+1)
	for i := 0; i < n+m+1; i++ {
		fmt.Fscan(in, &c[i])
	}

	b := make([]int, m+1)
	for i := m; i >= 0; i-- {
		b[i] = c[i+n] / a[n]
		for j := 0; j <= n; j++ {
			c[i+j] -= b[i] * a[j]
		}
	}
	for i := 0; i < m; i++ {
		fmt.Print(b[i], " ")
	}
	fmt.Println(b[m])
}
