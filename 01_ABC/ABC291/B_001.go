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
	X := make([]int, N*5)
	for i := 0; i < N*5; i++ {
		fmt.Fscan(in, &X[i])
	}
	sort.Ints(X)
	sum := 0.0
	for i := N; i < 4*N; i++ {
		sum += float64(X[i])
	}
	fmt.Println(sum / float64(3*N))
}
