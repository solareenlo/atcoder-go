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
	ans := make([]int, n)
	for i := range ans {
		ans[i] = m
	}
	for m > 0 {
		m--
		var v int
		fmt.Fscan(in, &v)
		v--
		ans[v]--
	}
	for i := range ans {
		fmt.Println(ans[i])
	}
}
