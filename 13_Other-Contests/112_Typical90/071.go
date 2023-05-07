package main

import (
	"bufio"
	"fmt"
	"os"
)

var N, K int
var G [100005][]int
var deg [100005]int
var s []int
var Ans [][]int
var v []int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var M int
	fmt.Fscan(in, &N, &M, &K)
	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		G[u] = append(G[u], v)
		deg[v]++
	}
	for i := 1; i <= N; i++ {
		if deg[i] == 0 {
			s = append(s, i)
		}
	}
	v = make([]int, 100005)
	for i := range v {
		v[i] = -1
	}
	Dfs(1)
	if len(Ans) != K {
		fmt.Fprintln(out, -1)
	} else {
		for _, a := range Ans {
			for i := 1; i <= N; i++ {
				if i == N {
					fmt.Fprintln(out, a[i])
				} else {
					fmt.Fprintf(out, "%d ", a[i])
				}
			}
		}
	}
}

func Dfs(d int) bool {
	if d == N+1 {
		tmp := make([]int, len(v))
		copy(tmp, v)
		Ans = append(Ans, tmp)
		return true
	}
	if len(s) == 0 {
		return false
	}
	for i := len(s) - 1; i >= 0; i-- {
		if len(Ans) == K {
			break
		}
		cur := s[i]
		s = erase(s, i)
		for _, nxt := range G[cur] {
			deg[nxt]--
			if deg[nxt] == 0 {
				s = append(s, nxt)
			}
		}
		v[d] = cur
		res := Dfs(d + 1)
		if !res {
			return false
		}
		for _, nxt := range G[cur] {
			if deg[nxt] == 0 {
				s = s[:len(s)-1]
			}
			deg[nxt]++
		}
		s = insert(&s, i, cur)
	}
	return true
}

func erase(a []int, pos int) []int {
	return append(a[:pos], a[pos+1:]...)
}

func insert(a *[]int, index int, value int) []int {
	n := len(*a)
	if index < 0 {
		index = (index%n + n) % n
	}
	switch {
	case index == n:
		return append(*a, value)

	case index < n:
		*a = append((*a)[:index+1], (*a)[index:]...)
		(*a)[index] = value
		return *a

	case index < cap(*a):
		*a = (*a)[:index+1]
		for i := n; i < index; i++ {
			(*a)[i] = 0
		}
		(*a)[index] = value
		return *a

	default:
		b := make([]int, index+1)
		if n > 0 {
			copy(b, *a)
		}
		b[index] = value
		return b
	}
}
