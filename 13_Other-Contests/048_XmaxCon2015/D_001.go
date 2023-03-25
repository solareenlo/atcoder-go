package main

import (
	"bufio"
	"fmt"
	"os"
)

const nmax = 500010

var (
	n     int
	a     [nmax]int
	s     string
	graph [nmax][]int
	Len   [nmax]int
	ckmp  [nmax][26]int
	kmp   [nmax]int
	ans   [nmax]int
)

func dfs(v, p, d int) {
	Len[v] = d
	tmp := -1
	if p != -1 {
		kmp[v] = ckmp[p][int(s[v-1]-'a')]
		ans[v] = Len[v] - Len[kmp[v]]
		for i := 0; i < 26; i++ {
			ckmp[v][i] = ckmp[kmp[v]][i]
		}
		if kmp[v] == p {
			ckmp[v][int(s[v-1]-'a')] = v
		}
		tmp = ckmp[p][int(s[v-1]-'a')]
		ckmp[p][int(s[v-1]-'a')] = v
	}
	for _, v2 := range graph[v] {
		dfs(v2, v, d+1)
	}
	if p != -1 {
		ckmp[p][int(s[v-1]-'a')] = tmp
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		graph[a[i]] = append(graph[a[i]], i)
	}
	fmt.Fscan(in, &s)
	dfs(0, -1, 0)
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, ans[i])
	}
}
