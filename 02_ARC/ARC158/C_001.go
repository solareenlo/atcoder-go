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
	a := make([]int, n+1)
	b := make([]int, n+1)
	A := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		x := a[i]
		for x != 0 {
			A += x % 10
			x /= 10
		}
	}
	A *= 2 * n
	for O := 10; O <= 1e16; O *= 10 {
		for i := 1; i <= n; i++ {
			b[i] = a[i] % O
		}
		sort.Ints(b[1:])
		for r, l := n, 1; r > 0; r-- {
			for l <= n && b[l]+b[r] < O {
				l++
			}
			A -= 9 * (n - l + 1)
		}
	}
	fmt.Println(A)
}
