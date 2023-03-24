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

	var N int
	fmt.Fscan(in, &N)
	for N != 0 {
		for i := 0; i < N; i++ {
			for j := 0; j <= i; j++ {
				if i == j && j != 0 && j < N-1 {
					fmt.Fprint(out, "@")
					out.Flush()
				} else if i < N-2 {
					fmt.Fprint(out, "o")
					out.Flush()
				} else if i == N-2 {
					fmt.Fprint(out, "@")
					out.Flush()
				} else {
					fmt.Fprint(out, "#")
					out.Flush()
				}
			}
			fmt.Fprintln(out)
			out.Flush()
		}
		s := make([]string, N)
		for i := 0; i < N; i++ {
			fmt.Fscan(in, &s[i])
		}
		edge := make([][]int, N+1)
		for i := 0; i < N; i++ {
			for j := 0; j <= i; j++ {
				if s[i][j] == '@' {
					edge[j+N-i] = append(edge[j+N-i], j+1)
					edge[j+1] = append(edge[j+1], j+N-i)
				}
			}
		}
		Q := make([]int, 0)
		ans := make([]int, 0)
		for i := 1; i <= N; i++ {
			if len(edge[i]) == 1 {
				Q = append(Q, i)
			}
		}
		used := make([]bool, N+1)
		used[Q[0]] = true
		for len(Q) > 0 {
			cn := Q[0]
			Q = Q[1:]
			ans = append(ans, cn)
			for _, i := range edge[cn] {
				if len(edge[i]) <= 3 {
					if used[i] {
						continue
					}
					used[i] = true
					Q = append(Q, i)
				}
			}
		}
		for i := 1; i <= N; i++ {
			if used[i] {
				continue
			}
			ans = append(ans, i)
		}
		ret := make([]int, N)
		for i := 0; i < N; i++ {
			ret[ans[i]-1] = i + 1
		}
		for i := 0; i < N; i++ {
			if i != 0 {
				fmt.Fprint(out, " ")
				out.Flush()
			}
			fmt.Fprint(out, ret[i])
			out.Flush()
		}
		fmt.Fprintln(out)
		out.Flush()
		fmt.Fscan(in, &N)
	}
}
