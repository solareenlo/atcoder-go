package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, D int
	fmt.Fscan(in, &N, &D)
	R := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &R[i])
	}
	sort.Ints(R)
	ans := 0
	l := 0
	for r := 0; r < N; r++ {
		for R[l]+D < R[r] {
			l++
		}
		ans += (r - l) * (r - l - 1) / 2
	}
	fmt.Println(ans)
}
