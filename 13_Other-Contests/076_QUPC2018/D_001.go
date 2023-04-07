package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, l int
	fmt.Fscan(in, &n, &m, &l)
	var t [1 << 17]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t[i])
	}
	A := make([][]int, 1<<17)
	B := make([][]int, 1<<17)
	M := 0
	for i := 0; i < m+l; i++ {
		var x, a int
		fmt.Fscan(in, &x, &a)
		if i < m {
			A[x-1] = append(A[x-1], a)
		} else {
			B[x-1] = append(B[x-1], a)
		}
		if i < m && M < a {
			M = a
		}
	}
	type pair struct {
		x, y int
	}
	ans := make([]pair, 0)
	for i := 0; i < n; i++ {
		sort.Ints(A[i])
		sort.Ints(B[i])
		f := -1
		for _, b := range B[i] {
			for f+1 < len(A[i]) && A[i][f+1]+t[i] < b {
				f++
			}
			if f >= 0 {
				ans = append(ans, pair{b + t[i], A[i][f]})
			}
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		if ans[i].x == ans[j].x {
			return ans[i].y < ans[j].y
		}
		return ans[i].x < ans[j].x
	})
	p := 0
	cnt := 0
	for _, q := range ans {
		if q.y > p {
			p = q.x
			cnt += 2
		}
	}
	if p < M {
		fmt.Println(cnt + 1)
	} else {
		fmt.Println(cnt)
	}
}
