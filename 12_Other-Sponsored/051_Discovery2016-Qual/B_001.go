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
	a := make([][]int, 100001)
	m := 0
	for i := 0; i < n; i++ {
		var b int
		fmt.Fscan(in, &b)
		a[b] = append(a[b], i)
		m = max(m, b)
	}
	s := 1
	b := 0
	for i := 0; i < m+1; i++ {
		d := 0
		for _, c := range a[i] {
			if c < b {
				c += n
			}
			d = max(d, c)
		}
		if len(a[i]) > 0 {
			if d >= n && (i != m || d != n) {
				s++
			}
			b = d % n
		}
	}
	fmt.Println(s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
