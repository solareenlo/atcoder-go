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

	p := make([]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		if a < b {
			p[b]++
		} else {
			p[a]++
		}
	}

	cnt := 0
	for i := 0; i < n; i++ {
		if p[i] == 1 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
