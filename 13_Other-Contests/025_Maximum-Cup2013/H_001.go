package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)
	const N = 2048

	var n, m int
	fmt.Fscan(in, &n, &m)
	var inputs [5000][4]int
	vs := make([]int, 0)
	for i := 0; i < m; i++ {
		var cmp string
		var a, b, c int
		fmt.Fscan(in, &a, &cmp, &b, &c)
		inputs[i][0] = a
		if cmp[0] == 's' {
			inputs[i][1] = 1
		} else {
			inputs[i][1] = 0
		}
		inputs[i][2] = b
		inputs[i][3] = c
		vs = append(vs, a)
		vs = append(vs, b)
	}
	sort.Ints(vs)
	vs = unique(vs)
	n = len(vs)
	var mat [N][N]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = INF
		}
	}
	for i := 0; i < m; i++ {
		inputs[i][0] = upperBound(vs, inputs[i][0]) - 1
		inputs[i][2] = upperBound(vs, inputs[i][2]) - 1
		inputs[i][3] = upperBound(vs, inputs[i][3]) - 1
		for j := inputs[i][2]; j <= inputs[i][3]; j++ {
			if inputs[i][1] == 0 {
				mat[inputs[i][0]][j] = min(mat[inputs[i][0]][j], i+1)
			} else {
				mat[j][inputs[i][0]] = min(mat[j][inputs[i][0]], i+1)
			}
		}
	}
	ans := 0
	low := 1
	high := m

	vcnt := 0
	v := make([]int, N)
	par := make([]int, N)
	var dfs func(int, int) bool
	dfs = func(nod, lim int) bool {
		par[nod] = 1
		v[nod] = vcnt
		for i := 0; i < n; i++ {
			if mat[nod][i] <= lim {
				if par[i] != 0 {
					par[nod] = 0
					return true
				}
				if v[i] == vcnt {
					continue
				}
				if dfs(i, lim) {
					par[nod] = 0
					return true
				}
			}
		}
		par[nod] = 0
		return false
	}

	var findCycle func(int) bool
	findCycle = func(val int) bool {
		vcnt++
		for i := 0; i < n; i++ {
			if v[i] != vcnt {
				if dfs(i, val) {
					return true
				}
			}
		}
		return false
	}
	for low <= high {
		mid := (low + high) / 2
		if findCycle(mid) {
			high = mid - 1
			ans = mid
		} else {
			low = mid + 1
		}
	}
	fmt.Println(ans)
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
