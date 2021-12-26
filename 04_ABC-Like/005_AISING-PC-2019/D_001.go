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

	var n, q int
	fmt.Fscan(in, &n, &q)

	a := make([]int, 100001)
	ac := make([]int, 100001)
	ac1 := make([]int, 100001)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		ac[i+1] = a[i] + ac[i]
		if (i+n)&1 != 0 {
			ac1[i/2+1] = a[i] + ac1[i/2]
		}
	}

	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(in, &x)
		l := 1
		r := (n+1)/2 + 1
		for r-l > 1 {
			m := (l + r) / 2
			if a[n+1-m*2]+a[n-m] < x*2 {
				r = m
			} else {
				l = m
			}
		}
		fmt.Fprintln(out, ac[n]-ac[n-l]+ac1[(n-l*2+1)/2])
	}

}
