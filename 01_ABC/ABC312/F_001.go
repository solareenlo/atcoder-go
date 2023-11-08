package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const M = 200005

	var a, s [5][]int
	var l [5]int
	for i := range a {
		a[i] = make([]int, M)
	}
	for i := range s {
		s[i] = make([]int, M)
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		var t, x int
		fmt.Fscan(in, &t, &x)
		t++
		l[t]++
		a[t][l[t]] = x
	}
	for i := 1; i <= 3; i++ {
		tmp := a[i][1 : l[i]+1]
		sort.Slice(tmp, func(a, b int) bool {
			return tmp[a] > tmp[b]
		})
	}
	for i := 1; i <= 3; i++ {
		for j := 1; j <= l[i]; j++ {
			s[i][j] = s[i][j-1] + a[i][j]
		}
	}
	ans := 0
	for i := 0; i <= l[2]; i++ {
		k := lowerBound(s[3][:+1+l[3]], i)
		if k+i > m || k > l[3] {
			break
		}
		ans = max(ans, s[2][i]+s[1][min(m-i-k, l[1])])
	}
	fmt.Println(ans)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
