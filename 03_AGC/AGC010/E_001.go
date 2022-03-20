package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 2005

var (
	vis = make([]bool, N)
	n   int
	d   = make([]int, N)
	a   = make([]int, N)
	g   = make([][]int, N)
)

func DFS(k1 int) {
	vis[k1] = true
	for i := 1; i <= n; i++ {
		if !vis[i] && gcd(a[k1], a[i]) > 1 {
			g[k1] = append(g[k1], i)
			d[i]++
			DFS(i)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	tmp := a[1 : n+1]
	sort.Ints(tmp)

	for i := 1; i <= n; i++ {
		if !vis[i] {
			DFS(i)
		}
	}

	for i := 1; i <= n; i++ {
		for j := n; j >= 1; j-- {
			if d[j] == 0 {
				fmt.Fprint(out, a[j], " ")
				for _, k := range g[j] {
					d[k]--
				}
				d[j] = -1
				break
			}
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
