package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, a, b int
	fmt.Fscan(in, &N)

	imo := make([]int, 100005)
	S := make([]struct{ a, b int }, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a, &b)
		imo[a]++
		imo[b]--
		S[i] = struct{ a, b int }{a, b}
	}

	for i := 1; i <= 100000; i++ {
		imo[i] += imo[i-1]
	}

	ma := -1
	L := -1
	R := -1

	for i := 0; i <= 100000; i++ {
		if ma < imo[i] {
			ma = imo[i]
			L = i
			R = i
		} else if ma == imo[i] {
			R = i
		}
	}

	for i := 0; i < N; i++ {
		if S[i].a <= L && R < S[i].b {
			fmt.Println(ma - 1)
			return
		}
	}

	fmt.Println(ma)
}
