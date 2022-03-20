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

	p := make([]int, n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i], &q[i])
	}
	sort.Ints(p)
	sort.Ints(q)

	a, b, c := 0, 0, 0
	for i := 0; i < n; i++ {
		a = p[i] + q[i]
		c += (a - b) * i * (n - i)
		b = a
	}
	fmt.Println(c)
}
