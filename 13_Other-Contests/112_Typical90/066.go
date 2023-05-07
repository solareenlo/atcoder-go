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

	l := make([]int, n)
	r := make([]int, n)
	s := 0.0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l[i], &r[i])
		for j := 0; j < i; j++ {
			a := 0
			for x := l[j]; x <= r[j]; x++ {
				a += min(x, r[i]+1) - min(x, l[i])
			}
			s += float64(a) / float64(r[i]-l[i]+1.0) / float64(r[j]-l[j]+1.0)
		}
	}
	fmt.Println(s)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
