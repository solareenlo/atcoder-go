package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100005

var (
	cn int
	cm int
	w  = [4]int{}
	c  = [N]int{}
	g  = make([][]int, N)
	h  = make([][]int, N)
	f  bool
)

func dfs(x int) {
	w[c[x]]++
	cn++
	for _, i := range g[x] {
		cm++
		if c[i] == 0 {
			c[i] = c[x]%3 + 1
			dfs(i)
		} else if c[i] != c[x]%3+1 {
			f = true
		}
	}
	for _, i := range h[x] {
		if c[i] == 0 {
			c[i] = (c[x]+1)%3 + 1
			dfs(i)
		} else if c[i] != (c[x]+1)%3+1 {
			f = true
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for k := 0; k < m; k++ {
		var i, j int
		fmt.Fscan(in, &i, &j)
		g[i] = append(g[i], j)
		h[j] = append(h[j], i)
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if c[i] == 0 {
			w[1] = 0
			w[2] = 0
			w[3] = 0
			cn = 0
			cm = 0
			f = false
			c[i] = 1
			dfs(i)
			if f {
				ans += cn * cn
			} else if w[2] == 0 || w[3] == 0 {
				ans += cm
			} else {
				ans += w[1]*w[2] + w[2]*w[3] + w[3]*w[1]
			}
		}
	}
	fmt.Println(ans)
}
