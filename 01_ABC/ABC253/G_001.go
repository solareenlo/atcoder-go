package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, b, d int
	fmt.Fscan(in, &n, &b, &d)

	cnt := n - 1
	a := 0
	for b > cnt {
		a++
		b -= cnt
		cnt--
	}
	a++
	b += a
	cnt = n - 1
	c := 0
	for d > cnt {
		c++
		d -= cnt
		cnt--
	}
	c++
	d += c
	const N = 2000005
	p := make([]int, N)
	for i := 1; i <= n; i++ {
		p[i] = i
	}
	for i := b; i <= n; i++ {
		p[a], p[i] = p[i], p[a]
	}
	tmp := p[a+1 : n+1]
	tmp = reverseOrderInt(tmp)
	for i := a + 1; i < n+1; i++ {
		p[i] = tmp[i-a-1]
	}
	tmp = p[c : n+1]
	tmp = reverseOrderInt(tmp)
	for i := c; i < n+1; i++ {
		p[i] = tmp[i-c]
	}
	for i := c + 1; i <= d; i++ {
		p[c], p[i] = p[i], p[c]
	}
	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d ", p[i])
	}
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
