package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m, ans int
var a [15]int

func dfs(last, s int) {
	if s == (1<<n)-1 {
		ans++
	}
	for i := last + 1; i <= m; i++ {
		dfs(i, s|a[i])
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var c int
		fmt.Fscan(in, &c)
		for ; c > 0; c-- {
			var x int
			fmt.Fscan(in, &x)
			a[i] |= 1 << (x - 1)
		}
	}
	dfs(0, 0)
	fmt.Println(ans)
}
