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

	var N, K int
	fmt.Fscan(in, &N, &K)
	v := make([]int, N+1)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &v[i+1])
		v[i+1] += v[i]
	}

	for i := K; i <= N; i++ {
		fmt.Fprintln(out, v[i]-v[i-K])
	}
}
