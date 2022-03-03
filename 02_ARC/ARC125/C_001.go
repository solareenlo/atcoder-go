package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s, m int
	fmt.Fscan(in, &s, &m)

	a := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &a[i])
	}

	vis := make([]bool, 200005)
	for i, j := 1, 1; i < m; i++ {
		fmt.Fprint(out, a[i], " ")
		vis[a[i]] = true
		for j < a[i] && vis[j] {
			j++
		}
		if !vis[j] {
			fmt.Fprint(out, j, " ")
			vis[j] = true
		}
	}
	for i := s; i > 0; i-- {
		if !vis[i] {
			fmt.Fprint(out, i, " ")
		}
	}
}
