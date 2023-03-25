package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	A := make([]map[int]bool, N)
	for i := 0; i < N; i++ {
		A[i] = make(map[int]bool)
	}

	for i := 0; i < M; i++ {
		var p, q int
		fmt.Fscan(in, &p, &q)
		p, q = p-1, q-1
		A[p][q] = true
		A[q][p] = true
	}

	S := make(map[int]bool)
	for i := 1; i < N; i++ {
		S[i] = true
	}

	B := []int{0}
	var ans []pair
	for i := 0; i < N; i++ {
		if i >= len(B) {
			fmt.Println("No")
			os.Exit(0)
		}
		for it := range S {
			if A[B[i]][it] {
				continue
			} else {
				B = append(B, it)
				ans = append(ans, pair{B[i], it})
				delete(S, it)
			}
		}
	}

	fmt.Println("Yes")
	w := bufio.NewWriter(os.Stdout)
	for _, p := range ans {
		fmt.Fprintf(w, "%d %d\n", p.first+1, p.second+1)
	}
	w.Flush()
}

type pair struct {
	first, second int
}
