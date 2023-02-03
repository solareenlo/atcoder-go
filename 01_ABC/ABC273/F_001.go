package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 3005

	var n, T int
	fmt.Fscan(in, &n, &T)
	var q [N]int
	x := make([]int, N)
	m := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &q[i])
		m++
		x[m] = q[i]
	}
	var s [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
		m++
		x[m] = s[i]
	}
	m++
	x[m] = 0
	m++
	x[m] = T
	tmp := x[1 : m+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	m = len(unique(x[1 : 1+m]))
	h := make(map[int]int)
	for i := 1; i <= m; i++ {
		h[x[i]] = i
	}
	var to [N]int
	var wall [N]bool
	for i := 1; i <= n; i++ {
		to[h[q[i]]] = h[s[i]]
		wall[h[q[i]]] = true
	}
	var f [N][N][2]int
	for i := range f {
		for j := range f[i] {
			for k := range f[i][j] {
				f[i][j][k] = int(1e18)
			}
		}
	}
	S := h[0]
	T = h[T]
	f[S][S][0] = 0
	f[S][S][1] = 0
	for len := 2; len <= m; len++ {
		for i, j := 1, 1+len-1; j <= m; i, j = i+1, j+1 {
			if !wall[i] || wall[i] && to[i] >= i && to[i] <= j {
				f[i][j][0] = min(f[i][j][0], f[i+1][j][0]+x[i+1]-x[i], f[i+1][j][1]+x[j]-x[i])
			}
			if !wall[j] || wall[j] && to[j] >= i && to[j] <= j {
				f[i][j][1] = min(f[i][j][1], f[i][j-1][1]+x[j]-x[j-1], f[i][j-1][0]+x[j]-x[i])
			}
		}
	}

	ans := int(1e18)
	for i := 1; i <= T; i++ {
		for j := T; j <= m; j++ {
			ans = min(ans, f[i][j][0], f[i][j][1])
		}
	}
	if ans == int(1e18) {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}

}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
