package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	c := make([]int, n)
	adj := make([][]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
		c[a[i]]++
	}
	ans := n
	for i := 0; i < n; i++ {
		if c[i] == 0 {
			ans--
			adj[0] = append(adj[0], i)
		}
	}
	for i := 0; i < n; i++ {
		if len(adj[i]) == 0 {
			break
		}
		for _, u := range adj[i] {
			c[a[u]]--
			if c[a[u]] == 0 {
				adj[i+1] = append(adj[i+1], a[u])
				ans--
			}
		}
	}
	fmt.Println(ans)
}
