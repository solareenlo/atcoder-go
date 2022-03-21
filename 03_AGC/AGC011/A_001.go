package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, c, k int
	fmt.Fscan(in, &n, &c, &k)

	t := make([]int, 100100)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &t[i])
	}
	tmp := t[1 : n+1]
	sort.Ints(tmp)

	nw := 0
	la := 0
	s := 0
	ans := 0
	for nw < n {
		nw++
		la = t[nw] + k
		s = c - 1
		for t[nw+1] <= la && s > 0 && nw < n {
			s--
			nw++
		}
		ans++
	}
	fmt.Println(ans)
}
