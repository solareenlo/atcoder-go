package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var a, b string
	fmt.Fscan(in, &n, &a, &b)
	a = " " + a
	b = " " + b

	ans := 0
	k := 0
	s := 0
	c := make([]int, 1000005)
	for i, j := n, n; i > 0; i-- {
		s += c[i+k]
		if j > i || b[i] != a[j] {
			for j != 0 && (j > i || b[i] != a[j]) {
				j--
			}
			if j == i {
				continue
			}
			if j == 0 {
				fmt.Println(-1)
				return
			}
			k++
			s++
			c[i-1+k] = 0
			c[j-1+k]--
		}
		if s > ans {
			ans = s
		}
	}
	fmt.Println(ans)
}
