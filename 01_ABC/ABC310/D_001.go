package main

import (
	"bufio"
	"fmt"
	"os"
)

var N, T, M int
var res int = 0
var A, B [45]int
var t [11]int

func f(i, m int) {
	if i > N {
		for j := 0; j < M; j++ {
			if t[A[j]] == t[B[j]] {
				return
			}
		}
		if m == T {
			res++
		}
		return
	}
	for j := 1; j <= m+1; j++ {
		t[i] = j
		if j <= m {
			f(i+1, m)
		} else {
			f(i+1, m+1)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &T, &M)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &A[i], &B[i])
	}
	f(1, 0)
	fmt.Println(res)
}
