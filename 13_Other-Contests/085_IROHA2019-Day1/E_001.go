package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, A, B int
	fmt.Fscan(in, &N, &A, &B)

	D := make([]int, B+2)
	for i := 0; i < B; i++ {
		fmt.Fscan(in, &D[i+1])
	}
	D[B+1] = N + 1
	sort.Ints(D)

	ans := N - B
	for i := 0; i < B+1; i++ {
		ans -= (D[i+1] - D[i] - 1) / A
	}
	fmt.Println(ans)
}
