package main

import (
	"bufio"
	"fmt"
	"os"
)

var p = int64(499903)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	out.Flush()

	var n, m int64
	fmt.Scan(&n, &m)

	root := make(map[int64]int64)
	for i := int64(0); i+i < p; i++ {
		root[(i*i)%p] = i
	}

	res := make([]int64, n)
	for i := int64(0); i < n; i++ {
		res[i] = (2 * p * i) + (i*i)%p
		fmt.Fprint(out, res[i])
		if i != n-1 {
			fmt.Fprint(out, " ")
		}
	}
	fmt.Fprint(out, "\n")

	var q, s, rem, wa, rt, k, a, b int64
	fmt.Scan(&q)
	for i := int64(0); i < q; i++ {
		fmt.Fscan(in, &s)
		rem = s % p
		wa = s / (2 * p)
		rt = root[(2*rem+p-((wa*wa)%p))%p]
		if (wa+rt)%2 == 0 {
			k = (wa + rt) / 2
		} else {
			k = (wa + rt + p) / 2
		}
		a = k
		b = wa - k
		if a < b {
			fmt.Fprint(out, a+1, " ", b+1)
		} else {
			fmt.Fprint(out, b+1, " ", a+1)
		}
	}
}
