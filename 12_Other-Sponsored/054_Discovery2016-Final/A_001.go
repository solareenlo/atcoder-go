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
	var a [514]int
	var x [114]int = [114]int{211, 100000, 50000, 30000, 20000, 10000}
	ans := 0
	for i := 1; i <= n; i++ {
		var c int
		fmt.Fscan(in, &c)
		a[c] = x[i]
		if ans < c {
			ans = c
		}
	}
	for i := 1; i <= ans; i++ {
		fmt.Println(a[i])
	}
}
