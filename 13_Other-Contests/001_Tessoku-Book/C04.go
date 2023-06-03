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
	A := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			A = append(A, i)
			A = append(A, n/i)
		}
	}
	sort.Ints(A)
	for i := 0; i < len(A); i++ {
		fmt.Println(A[i])
	}
}
