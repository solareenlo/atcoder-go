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

	h := make([]int, n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &h[i])
	}

	for i := 0; i < n; i++ {
		p := i - 1
		for p != -1 && h[p] <= h[i] {
			p = a[p]
		}
		a[i] = p
	}
	for i := n - 1; i >= 0; i-- {
		p := i + 1
		for p != n && h[p] <= h[i] {
			p = b[p]
		}
		b[i] = p
	}
	for i := 0; i < n; i++ {
		fmt.Println(b[i] - a[i] - 2)
	}
}
