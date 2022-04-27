package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)
	m := make(map[string]int)
	for i := 0; i < N; i++ {
		var S string
		fmt.Fscan(in, &S)
		m[S]++
	}

	p := make([]int, N)
	q := make([]string, N)
	for S, a := range m {
		a = N - a
		p[a]++
		q[a] = S
	}

	X := 0
	for i := 0; i < N; i++ {
		X += p[i]
		if X == K && p[i] == 1 {
			fmt.Println(q[i])
			return
		}
	}
	fmt.Println("AMBIGUOUS")
}
