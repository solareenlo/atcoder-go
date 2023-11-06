package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var e [5][5]int

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n+1)
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		b[i] = a[i]
	}
	sort.Ints(b[1:])
	for i := 1; i <= n; i++ {
		e[a[i]][b[i]]++
	}
	P, O := 0, 0
	for i := 1; i <= 4; i++ {
		for j := i + 1; j <= 4; j++ {
			P += max(e[i][j], e[j][i])
			O = max(O, abs(e[i][j]-e[j][i]))
		}
	}
	fmt.Println(P - O)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
