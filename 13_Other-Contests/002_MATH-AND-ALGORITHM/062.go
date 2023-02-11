package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
		p[i]--
	}

	z := 0
	q := make([]int, n)
	for k > 0 {
		if k&1 != 0 {
			z = p[z]
		}
		for i := 0; i < n; i++ {
			q[i] = p[p[i]]
		}
		p, q = q, p
		k >>= 1
	}
	fmt.Println(z + 1)
}
