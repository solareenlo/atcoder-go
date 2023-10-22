package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF = int(1e9)
const N = int(1e5)

var n, m, q int
var g [N][]int
var qu, qv [N]int

func small() {
	var d [1000][1000]int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			d[i][j] = 1
		}
		d[i][i] = 0
	}
	for i := 0; i < n; i++ {
		for _, j := range g[i] {
			d[i][j] = INF
		}
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
	for i := 0; i < q; i++ {
		if d[qu[i]][qv[i]] == INF {
			fmt.Println(-1)
		} else {
			fmt.Println(d[qu[i]][qv[i]])
		}
	}
}

func large() {
	qq := make([][]int, n)
	ans := make([]int, q)
	for i := 0; i < q; i++ {
		u := qu[i]
		v := qv[i]
		if n-len(g[u]) > (n+1)/2 && n-len(g[v]) > (n+1)/2 {
			if binarySearch(g[u], 0, len(g[u])-1, v) != -1 {
				ans[i] = 2
			} else {
				ans[i] = 1
			}
		} else {
			if len(g[u]) < len(g[v]) {
				qu[i], qv[i] = qv[i], qu[i]
			}
			qq[qu[i]] = append(qq[qu[i]], i)
		}
	}
	dist := make([]int, n)
	for k := 0; k < n; k++ {
		if len(qq[k]) == 0 {
			continue
		}
		for i := range dist {
			dist[i] = -1
		}
		dist[k] = 0
		q := make([]int, 0)
		q = append(q, k)
		remain := make([]int, 0)
		for i := 0; i < n; i++ {
			if i != k {
				remain = append(remain, i)
			}
		}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			tmp := make([]int, 0)
			j := 0
			for i := 0; i < len(remain); i++ {
				for j < len(g[u]) && g[u][j] < remain[i] {
					j++
				}
				if j < len(g[u]) && g[u][j] == remain[i] {
					tmp = append(tmp, remain[i])
				} else {
					dist[remain[i]] = dist[u] + 1
					q = append(q, remain[i])
				}
			}
			remain, tmp = tmp, remain
		}
		for _, id := range qq[k] {
			ans[id] = dist[qv[id]]
		}
	}
	for i := 0; i < q; i++ {
		fmt.Println(ans[i])
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m, &q)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	for i := 0; i < n; i++ {
		sort.Ints(g[i])
	}
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &qu[i], &qv[i])
		qu[i]--
		qv[i]--
	}
	if n < 800 {
		small()
	} else {
		large()
	}
}

func binarySearch(nums []int, left, right, toFind int) int {
	if right >= left {
		mid := left + (right-left)/2
		if nums[mid] == toFind {
			return mid
		}
		if nums[mid] > toFind {
			return binarySearch(nums, left, mid-1, toFind)
		}
		return binarySearch(nums, mid+1, right, toFind)
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
