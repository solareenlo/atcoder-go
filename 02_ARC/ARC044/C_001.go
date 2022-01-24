package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	w := [2]int{}
	var q int
	fmt.Fscan(in, &w[0], &w[1], &q)

	const MX = 100004
	m := make([][][]int, 2)
	for i := range m {
		m[i] = make([][]int, MX)
	}
	for i := 0; i < q; i++ {
		var t, d, a int
		fmt.Fscan(in, &t, &d, &a)
		a--
		m[d][t] = append(m[d][t], a)
	}

	ans := 0
	for i := 0; i < 2; i++ {
		x := make([]int, MX)
		mini := 1 << 60
		for j := MX - 1; j >= 0; j-- {
			sz := len(m[i][j])
			b := 0
			sort.Ints(m[i][j])
			if sz == w[i] {
				fmt.Println(-1)
				return
			}
			for k := 0; k < sz; k++ {
				if k == sz-1 || m[i][j][k]+1 != m[i][j][k+1] {
					s := k - b + 1
					bn := m[i][j][b]
					pb := 1 << 60
					if bn != 0 {
						pb = x[bn-1] + 1
					}
					ae := x[bn+s] + s
					if bn+s == w[i] {
						ae = 1 << 60
					}
					for j := 0; j < s; j++ {
						x[bn+j] = min(ae-j, pb+j)
					}
					b = k + 1
				}
			}
		}

		for j := 0; j < w[i]; j++ {
			mini = min(mini, x[j])
		}
		ans += mini
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
