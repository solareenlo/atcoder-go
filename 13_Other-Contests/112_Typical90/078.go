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

	var d [100005]int
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		if u < v {
			d[v]++
		} else if u > v {
			d[u]++
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if d[i] == 1 {
			ans++
		}
	}
	fmt.Println(ans)
}
