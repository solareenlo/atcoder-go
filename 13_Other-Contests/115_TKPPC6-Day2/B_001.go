package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	var S, T string
	fmt.Fscan(in, &N, &S, &T)
	z := make([]int, 0)
	o := make([]int, 0)
	s := 0
	for i := 0; i < N; i++ {
		if S[i] != T[i] {
			if s == 0 {
				z = append(z, i)
			} else {
				o = append(o, i)
			}
		}
		if i != N-1 && S[i] == S[i+1] {
			s ^= 1
		}
	}

	if len(z) != len(o) {
		fmt.Println(-1)
	} else {
		res := 0
		for i := 0; i < len(o); i++ {
			res += abs(z[i] - o[i])
		}
		fmt.Println(res)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
