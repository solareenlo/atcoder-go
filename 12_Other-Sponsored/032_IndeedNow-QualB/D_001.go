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
	var res, lst [100005]int
	for i := range lst {
		lst[i] = -1
	}
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		x--
		res[x] += (i - lst[x]) * (i - lst[x] - 1) / 2
		lst[x] = i
	}
	for i := 0; i < m; i++ {
		res[i] += (n - lst[i]) * (n - lst[i] - 1) / 2
	}
	for i := 0; i < m; i++ {
		fmt.Println(n*(n+1)/2 - res[i])
	}
}
