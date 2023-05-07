package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	sort.Ints(x)
	sort.Ints(y)

	ans := 0
	for i := 0; i < N; i++ {
		ans += abs(x[i] - x[N/2])
	}
	for i := 0; i < N; i++ {
		ans += abs(y[i] - y[N/2])
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
