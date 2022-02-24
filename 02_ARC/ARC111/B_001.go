package main

import (
	"bufio"
	"fmt"
	"os"
)

var fa = [400400]int{}

func dfs(x int) int {
	if fa[x] != 0 {
		fa[x] = dfs(fa[x])
		return fa[x]
	} else {
		return x
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	ans := 0
	cab := make([]int, 400400)
	for i := 1; i <= n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a = dfs(a)
		b = dfs(b)
		if (a^b) != 0 && (cab[a] == 0 || cab[b] == 0) {
			fa[a] = b
			cab[b] |= cab[a]
			ans++
		} else if cab[a] == 0 {
			cab[a] = 1
			ans++
		}
	}

	fmt.Println(ans)
}
