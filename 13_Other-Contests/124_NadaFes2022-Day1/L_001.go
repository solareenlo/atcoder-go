package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var G, H, vs, chs [2 << 17][]int
var dep, ch, idx [2 << 17]int
var ans []int

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	for i := 1; i < N; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
	}
	dfs1(0, -1, 0)
	for d := N - 1; d >= 0; d-- {
		for _, u := range vs[d] {
			for _, v := range H[u] {
				chs[u] = append(chs[u], idx[v])
			}
			sort.Ints(chs[u])
		}
		sort.Slice(vs[d], func(l, r int) bool {
			return lessThanSlices(chs[vs[d][l]], chs[vs[d][r]])
		})
		id := 0
		for i := 0; i < len(vs[d]); i++ {
			if i > 0 && lessThanSlices(chs[vs[d][i-1]], chs[vs[d][i]]) {
				id++
			}
			idx[vs[d][i]] = id
		}
	}
	dfs2(0)
	for i := 0; i < len(ans); i++ {
		if i+1 == len(ans) {
			fmt.Println(ans[i])
		} else {
			fmt.Printf("%d ", ans[i])
		}
	}
}

func lessThanSlices(s1, s2 []int) bool {
	minLength := len(s1)
	if len(s2) < minLength {
		minLength = len(s2)
	}

	for i := 0; i < minLength; i++ {
		if s1[i] < s2[i] {
			return true // s1がs2よりも小さい
		} else if s1[i] > s2[i] {
			return false // s1がs2よりも大きい
		}
	}

	if len(s1) == len(s2) {
		return false // スライスは等しい
	} else if len(s1) < len(s2) {
		return true // s1がs2よりも小さい
	}

	return false // s1がs2よりも大きい
}

func dfs1(u, p, d int) int {
	dep[u] = d
	vs[d] = append(vs[d], u)
	ch[u] = 1
	for _, v := range G[u] {
		if v != p {
			H[u] = append(H[u], v)
			ch[u] += dfs1(v, u, d+1)
		}
	}
	return ch[u]
}

func dfs2(u int) {
	ans = append(ans, dep[u])
	sort.Slice(H[u], func(l, r int) bool {
		return idx[H[u][l]] < idx[H[u][r]]
	})
	for _, v := range H[u] {
		dfs2(v)
		ans = append(ans, dep[u])
	}
}
