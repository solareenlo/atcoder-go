package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, Q int
	fmt.Fscan(in, &N, &Q)
	W := make([]int, 1<<17)
	T := make([]int, 1<<17)
	x := make([]int, 1<<17)
	ans := make([]int, 1<<17)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &W[i], &T[i], &x[i])
		ans[i] = -1
	}
	sum := make([][]int, 4)
	for i := range sum {
		sum[i] = make([]int, 4000000)
	}
	sum[2][1] = 1
	id := 0
	for i := 0; i <= 3652425; i++ {
		for id < N && W[id] <= sum[3][i] {
			ans[id] = i
			if T[id] == 0 {
				sum[2][i+1]++
				sum[2][i+x[id]+1]--
			} else if T[id] == 1 {
				sum[1][i+1]++
				sum[1][i+x[id]+1]--
				sum[2][i+x[id]+1] -= x[id]
			} else {
				sum[0][i+1]++
				sum[0][i+2]++
				sum[0][i+x[id]+1] -= 2
				sum[1][i+x[id]+1] -= 2*x[id] - 1
				sum[2][i+x[id]+1] -= x[id] * x[id]
			}
			id++
		}
		sum[0][i+1] += sum[0][i]
		sum[1][i+1] += sum[1][i] + sum[0][i+1]
		sum[2][i+1] += sum[2][i] + sum[1][i+1]
		sum[3][i+1] += sum[3][i] + sum[2][i+1]
	}
	for i := 0; i < N; i++ {
		if ans[i] >= 0 {
			fmt.Fprintln(out, ans[i])
		} else {
			fmt.Fprintln(out, "Many years later")
		}
	}
	for i := 0; i < Q; i++ {
		var y int
		fmt.Fscan(in, &y)
		fmt.Fprintln(out, sum[3][y])
	}
}
