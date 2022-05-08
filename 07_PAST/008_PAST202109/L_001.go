package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)

	A := make([]int, N)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}

	S := make([]int, N+1)
	for i := 0; i < N; i++ {
		S[i+1] += S[i] + A[i]
	}
	sort.Ints(S)

	hi := 1 << 60
	lo := -1
	for hi-lo > 1 {
		mid := (hi + lo) / 2
		sum := 0
		pos := 0
		for i := 0; i < N+1; i++ {
			for S[i]-S[pos] > mid {
				pos++
			}
			sum += i - pos
		}
		if sum >= K {
			hi = mid
		} else {
			lo = mid
		}
	}
	fmt.Println(hi)
}
