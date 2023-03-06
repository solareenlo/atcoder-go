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

	var r, c, m, n int
	fmt.Fscan(in, &r, &c, &m, &n)

	ra := make([]int, n)
	rb := make([]int, n)
	ca := make([]int, n)
	cb := make([]int, n)
	s := make([][]int, r)
	for i := range s {
		s[i] = make([]int, c)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ra[i], &rb[i], &ca[i], &cb[i])
		ra[i]--
		rb[i]--
		ca[i]--
		cb[i]--
		for j := ra[i]; j <= rb[i]; j++ {
			for k := ca[i]; k <= cb[i]; k++ {
				s[j][k]++
			}
		}
	}
	for i := 0; i < n; i++ {
		res := 0
		for j := 0; j < r; j++ {
			for k := 0; k < c; k++ {
				if ra[i] <= j && j <= rb[i] && ca[i] <= k && k <= cb[i] {
					if s[j][k]%4 == 1 {
						res++
					}
				} else if s[j][k]%4 == 0 {
					res++
				}
			}
		}
		if res == m {
			fmt.Fprintln(out, i+1)
		}
	}
}
