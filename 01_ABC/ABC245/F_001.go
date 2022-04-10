package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	e := make([][]int, n)
	out := make([]int, n)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		e[y-1] = append(e[y-1], x-1)
		out[x-1]++
	}

	que := make([]int, 0)
	ans := n
	for i := 0; i < n; i++ {
		if out[i] == 0 {
			que = append(que, i)
		}
	}
	for len(que) > 0 {
		ans--
		cur := que[0]
		que = que[1:]
		sz := len(e[cur])
		for i := 0; i < sz; i++ {
			out[e[cur][i]]--
			if out[e[cur][i]] == 0 {
				que = append(que, e[cur][i])
			}
		}
	}
	fmt.Println(ans)
}
