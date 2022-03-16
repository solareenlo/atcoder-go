package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	const N = 200005
	p := [N][18]int{}
	s := [N][18]int{}
	ps := [N][18]int{}
	ss := [N][18]int{}
	for i := 1; i <= n; i++ {
		ps[i][0] = upperBound(a[1:], a[i]-k)
		p[i][0] = ps[i][0]
		ss[i][0] = lowerBound(a[1:], a[i]+k) + 1
		s[i][0] = ss[i][0]
	}

	for j := 1; j <= 17; j++ {
		for i := 1; i <= n; i++ {
			if p[i][j-1] >= 0 {
				p[i][j] = p[p[i][j-1]][j-1]
				s[i][j] = s[s[i][j-1]][j-1]
				ps[i][j] = ps[i][j-1] + ps[p[i][j-1]][j-1]
				ss[i][j] = ss[i][j-1] + ss[s[i][j-1]][j-1]
			}
		}
	}

	var q int
	fmt.Fscan(in, &q)
	for j := 0; j < q; j++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		x := l
		y := r
		ans := 0
		for i := 17; i >= 0; i-- {
			if s[x][i] != 0 && s[x][i] <= r {
				ans -= ss[x][i]
				ans += ps[y][i]
				ans += (1 << i)
				x = s[x][i]
				y = p[y][i]
			}
		}
		fmt.Fprintln(out, ans+r-l+1)
	}
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
