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

	var p string
	fmt.Fscan(in, &p)
	lenP := len(p)
	p += " "
	var w [201000]int
	var i int
	for i = 0; i < lenP; i++ {
		if p[i] != p[i+1] {
			w[i] = 1
		} else {
			w[i] = 0
		}
	}
	n := i
	var C [410][410][2]int
	for i := 0; i < n-1; i++ {
		for j := 1; j <= 400; j++ {
			C[j][i%j][w[i]]++
		}
	}
	var Q int
	fmt.Fscan(in, &Q)
	var tw [201000]int
	for Q > 0 {
		Q--
		var q string
		fmt.Fscan(in, &q)
		m := len(q)
		for i = 0; i < m; i++ {
			if q[i%m] != q[(i+1)%m] {
				tw[i] = 1
			} else {
				tw[i] = 0
			}
		}
		s := 0
		if m <= 400 {
			for i = 0; i < m; i++ {
				if i >= n {
					continue
				}
				if tw[i] == 0 {
					s += C[m][i][1]
				} else {
					s += C[m][i][0]
				}
			}
		} else {
			for i = m; i < n; i++ {
				tw[i] = tw[i-m]
			}
			for i = 0; i < n-1; i++ {
				if w[i] != tw[i] {
					s++
				}
			}
		}
		if p[0] != q[0] {
			s++
		}
		fmt.Fprintln(out, (s+1)/2)
	}
}
