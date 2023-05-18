package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var G [1 << 17][]int
var D, in [1 << 17]int
var P [17][1 << 17]int
var id int

func main() {
	IN := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(IN, &n, &q)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(IN, &u, &v)
		u--
		v--
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}
	dfs(0)
	for k := 0; k < 16; k++ {
		for i := 0; i < n; i++ {
			P[k+1][i] = P[k][P[k][i]]
		}
	}
	for q > 0 {
		q--
		var k int
		fmt.Fscan(IN, &k)
		A := make([]int, k)
		for i := 0; i < k; i++ {
			fmt.Fscan(IN, &A[i])
			A[i]--
		}
		sort.Slice(A, func(i, j int) bool {
			return in[A[i]] < in[A[j]]
		})
		ans := 0
		for i := 0; i < k; i++ {
			ans += dist(A[i], A[(i+1)%k])
		}
		u := 0
		v := 0
		for i := 0; i < k; i++ {
			if dist(A[u], A[v]) < dist(A[u], A[i]) {
				v = i
			}
		}
		for i := 0; i < k; i++ {
			if dist(A[v], A[u]) < dist(A[v], A[i]) {
				u = i
			}
		}
		ans -= dist(A[u], A[v])
		fmt.Println(ans)
	}
}

func dfs(u int) {
	in[u] = id
	id++
	for _, v := range G[u] {
		if v != P[0][u] {
			P[0][v] = u
			D[v] = D[u] + 1
			dfs(v)
		}
	}
}

func dist(u, v int) int {
	res := D[u] + D[v]
	if D[u] < D[v] {
		u, v = v, u
	}
	d := D[u] - D[v]
	for k := 0; k < 17; k++ {
		if (d >> k & 1) != 0 {
			u = P[k][u]
		}
	}
	if u != v {
		for k := 16; k >= 0; k-- {
			if P[k][u] != P[k][v] {
				u = P[k][u]
				v = P[k][v]
			}
		}
		u = P[0][u]
	}
	res -= 2 * D[u]
	return res
}
