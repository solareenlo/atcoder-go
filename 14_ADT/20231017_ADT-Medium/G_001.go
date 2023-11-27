package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200020

var vis map[string]bool
var n, m, w int
var v [N]int
var s [N]string

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
		w += len(s[i])
	}
	vis = make(map[string]bool)
	for m > 0 {
		var t string
		fmt.Fscan(in, &t)
		vis[t] = true
		m--
	}
	dfs(1, "", 0)
	fmt.Println(-1)
}

func dfs(x int, c string, b int) {
	if len(c) > 16 {
		return
	}
	if x > n {
		_, ok := vis[c]
		if len(c) > 2 && !ok {
			fmt.Println(c)
			os.Exit(0)
		}
		return
	}
	for i := 1; i <= n; i++ {
		if v[i] == 0 {
			if x < n {
				v[i] = 1
				t := c + s[i]
				for l := 1; l <= 16-w-b; l++ {
					t += "_"
					dfs(x+1, t, b+l)
				}
				v[i] = 0
			} else {
				dfs(x+1, c+s[i], b)
			}
		}
	}
}
