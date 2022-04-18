package main

import (
	"bufio"
	"fmt"
	"os"
)

var par = [200200]int{}

func find(x int) int {
	if x != par[x] {
		par[x] = find(par[x])
	}
	return par[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	for i := 0; i < N+M; i++ {
		par[i] = i
	}
	for i := 0; i < N; i++ {
		var x int
		fmt.Fscan(in, &x)
		for x > 0 {
			var y int
			fmt.Fscan(in, &y)
			y = y - 1 + N
			par[find(i)] = find(y)
			x--
		}
	}

	for i := 0; i < N; i++ {
		if find(i) != find(0) {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
	return
}
