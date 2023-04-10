package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const SIZE = 505
	const INF = 1000000005

	type P struct {
		x, y int
	}

	var C, A, ng [SIZE][SIZE]int

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
			C[i][j] = min(C[i][j], INF)
		}
		for j := 0; j <= i; j++ {
			fmt.Fscan(in, &A[i][j])
		}
	}
	vx := make([]P, 0)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if A[i][j] != 0 {
				if A[i][j] < C[i][j] {
					fmt.Println("IMPOSSIBLE")
					return
				} else if A[i][j] > C[i][j] {
					vx = append(vx, P{i, j})
				} else {
					ng[i][j] = 1
				}
			}
		}
	}
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			ng[i][j] |= ng[i+1][j]
			ng[i][j] |= ng[i+1][j+1]
		}
	}
	ans := P{-1, -1}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if ng[i][j] != 0 {
				continue
			}
			if len(vx) != 0 && C[i][j] >= INF {
				continue
			}
			up := true
			x := -1
			for k := 0; k < len(vx); k++ {
				p := vx[k]
				if p.x < i || p.y-j > p.x-i || p.y < j {
					up = false
					break
				}
				d := A[p.x][p.y] - C[p.x][p.y]
				if d%C[p.x-i][p.y-j] != 0 {
					up = false
					break
				}
				e := d / C[p.x-i][p.y-j]
				if x != -1 && x != e {
					up = false
					break
				}
				x = e
			}
			if up {
				if ans.x != -1 && ans.y != -1 {
					fmt.Println("AMBIGUOUS")
					return
				}
				ans = P{i, j}
			}
		}
	}
	if ans.x == -1 && ans.y == -1 {
		fmt.Println("IMPOSSIBLE")
	} else {
		fmt.Println(ans.x, ans.y)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
