package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, t int
	fmt.Fscan(in, &n, &t)
	su := 0
	dif := make([]int, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		su += a
		dif[i] = b - a
	}
	sort.Ints(dif)
	ans := 0
	for ans < n && su > t {
		su += dif[ans]
		ans++
	}
	if ans == n && su > t {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
