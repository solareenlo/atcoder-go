package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200010

	type Node struct {
		w, s, v int
	}

	var n int
	fmt.Fscan(in, &n)

	z := make([]Node, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &z[i].w, &z[i].s, &z[i].v)
	}
	tmp := z[1:]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].w+tmp[i].s < tmp[j].w+tmp[j].s
	})
	var f [N]int
	for i := 1; i <= n; i++ {
		for j := z[i].w + z[i].s; j >= z[i].w; j-- {
			f[j] = max(f[j], f[j-z[i].w]+z[i].v)
		}
	}
	res := 0
	for i := 0; i < N; i++ {
		res = max(res, f[i])
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
