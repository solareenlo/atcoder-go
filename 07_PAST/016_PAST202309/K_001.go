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

	const N = 1500

	type pair struct {
		x, y int
	}

	var S [N]string
	var A [N][N + 1]int
	var B [N + 1][N]int
	var D [N][N]int

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &S[i])
	}
	var sx, sy int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if S[i][j] == 'S' {
				sx = i
				sy = j
			}
			if S[i][j] == 'X' {
				A[i][j+1] = 1
				B[i+1][j] = 1
			}
			A[i][j+1] += A[i][j]
			B[i+1][j] += B[i][j]
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			D[i][j] = -1
		}
	}
	for k := 1; k < n; k++ {
		Q := make([]pair, 0)
		Q = append(Q, pair{sx, sy})
		D[sx][sy] = 0
		ans := -1
		for i := 0; i < len(Q); i++ {
			x, y := Q[i].x, Q[i].y
			if S[x][y] == 'G' {
				ans = D[x][y]
				break
			}
			for d := 0; d < 4; d++ {
				nx := x + (d-1)%2*k
				ny := y + (d-2)%2*k
				if nx < 0 || nx >= n || ny < 0 || ny >= n {
					continue
				}
				if S[nx][ny] == 'X' {
					continue
				}
				if x == nx && A[x][ny] != A[x][y] {
					continue
				}
				if y == ny && B[nx][y] != B[x][y] {
					continue
				}
				if ^D[nx][ny] != 0 {
					continue
				}
				D[nx][ny] = D[x][y] + 1
				Q = append(Q, pair{nx, ny})
			}
		}
		fmt.Fprintln(out, ans)
		for i := range Q {
			D[Q[i].x][Q[i].y] = -1
		}
	}
}
