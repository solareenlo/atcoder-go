package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	score := make([][2]int, N+1)
	for i := 1; i <= N; i++ {
		var c, p int
		fmt.Fscan(in, &c, &p)
		score[i][c-1] = p
	}

	for i := 1; i <= N; i++ {
		score[i][0] += score[i-1][0]
		score[i][1] += score[i-1][1]
	}

	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var L, R int
		fmt.Fscan(in, &L, &R)
		fmt.Println(score[R][0]-score[L-1][0], score[R][1]-score[L-1][1])
	}
}
