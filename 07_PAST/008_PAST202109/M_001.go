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

	var N, M int
	fmt.Fscan(in, &N, &M)

	type pair struct{ x, y int }
	G := make([][]pair, N)
	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		G[a] = append(G[a], pair{b, c})
		G[b] = append(G[b], pair{a, c})
	}

	visited := make([]bool, N)
	S := make([]int, 0)
	A := make([]int, N)
	B := make([]int, N)
	xmin := 0
	xmax := 1 << 60
	S = append(S, 0)
	A[0] = 1
	B[0] = 0
	for len(S) > 0 {
		n := S[0]
		S = S[1:]
		if visited[n] {
			continue
		}
		visited[n] = true
		for i := range G[n] {
			nxt := G[n][i].x
			c := G[n][i].y
			S = append(S, nxt)
			if A[n] == 1 {
				xmin = max(xmin, -B[n])
			} else {
				xmax = min(xmax, B[n])
			}
			a := -A[n]
			b := c - B[n]
			if a == 1 {
				xmin = max(xmin, -b)
			} else {
				xmax = min(xmax, b)
			}
			if A[nxt] == 0 && B[nxt] == 0 {
				A[nxt] = a
				B[nxt] = b
			} else if A[nxt] == a {
				if B[nxt] != b {
					xmin = xmax + 1
				}
			} else {
				if abs(b-B[nxt])%2 != 0 {
					xmin = xmax + 1
				} else {
					xmin = max(xmin, abs(b-B[nxt])/2)
					xmax = min(xmax, abs(b-B[nxt])/2)
				}
			}
		}
	}

	if xmin > xmax {
		fmt.Fprintln(out, -1)
	} else {
		for i := 0; i < N; i++ {
			fmt.Fprintln(out, A[i]*xmin+B[i])
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
