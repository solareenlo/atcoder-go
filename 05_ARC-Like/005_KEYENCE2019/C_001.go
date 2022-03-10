package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	s := 0
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
		a[i] -= b[i]
		s += a[i]
	}

	if s < 0 {
		fmt.Println(-1)
		return
	}

	tmp := a[1 : n+1]
	sort.Ints(tmp)
	s = 0
	j := 0
	for i := 1; a[i] < 0; i++ {
		s += a[i]
		j++
	}
	i := n
	for ; s < 0; i-- {
		s += a[i]
	}
	fmt.Println(n - i + j)
}
