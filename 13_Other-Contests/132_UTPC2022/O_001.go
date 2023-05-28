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
	var p, a [1000010]int
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}
	var m int
	fmt.Fscan(in, &m)
	for i := 0; i < m; i++ {
		var x int
		fmt.Fscan(in, &x)
		a[x]++
	}
	ans := 0
	for i := n; i > 1; i-- {
		if a[i] > 1 {
			ans += a[i] - 1
			a[p[i]] += a[i] - 1
		}
	}
	fmt.Println(ans)
}
