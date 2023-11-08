package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, K int
	fmt.Scan(&N, &K)

	ans := make([]int, N+1)
	r := 0
	for i := 1; i <= K+1; i++ {
		query := "? "
		for j := 1; j <= K+1; j++ {
			if i == j {
				continue
			}
			query += strconv.Itoa(j) + " "
		}
		fmt.Fprintln(out, query)
		out.Flush()
		fmt.Scan(&ans[i])
		if ans[i] == -1 {
			return
		}
		r ^= ans[i]
	}
	for i := 1; i <= K+1; i++ {
		ans[i] ^= r
	}

	query := "? "
	s := 0
	for i := 1; i < K; i++ {
		query += strconv.Itoa(i) + " "
		s ^= ans[i]
	}
	for i := K + 2; i < N+1; i++ {
		fmt.Fprintln(out, query+strconv.Itoa(i))
		out.Flush()
		var t int
		fmt.Scan(&t)
		if t == -1 {
			return
		}
		ans[i] = s ^ t
	}

	fmt.Fprintf(out, "! ")
	out.Flush()
	for i := 1; i <= N; i++ {
		fmt.Fprintf(out, "%d ", ans[i])
		out.Flush()
	}
	fmt.Fprintln(out)
	out.Flush()
}
