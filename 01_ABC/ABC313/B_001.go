package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var p [105]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		p[v]++
	}
	t, k := 0, 1
	for i := 1; i <= n; i++ {
		if p[i] == 0 {
			t++
			k = i
		}
	}
	if t == 1 {
		fmt.Println(k)
	} else {
		fmt.Println(-1)
	}
}
