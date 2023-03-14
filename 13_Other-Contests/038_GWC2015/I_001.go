package main

import (
	"bufio"
	"fmt"
	"os"
)

const SIZE = 505

var n, q, X, A, B, C int
var H [SIZE][SIZE]int
var dp [SIZE][SIZE]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &q)
	fmt.Fscan(in, &X, &A, &B, &C)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &H[i][j])
		}
	}
	for i := 0; i < q; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				dp[j][k] = H[j][k]
				if j == 0 || j == n-1 || k == 0 || k == n-1 {
					dp[j][k] = 1
				}
			}
		}
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if j > 0 {
					dp[j][k] = min(dp[j][k], dp[j-1][k]+1)
				}
				if k > 0 {
					dp[j][k] = min(dp[j][k], dp[j][k-1]+1)
				}
				if j > 0 && k > 0 {
					dp[j][k] = min(dp[j][k], dp[j-1][k-1]+1)
				}
			}
		}
		for j := n - 1; j >= 0; j-- {
			for k := 0; k < n; k++ {
				if j+1 < n {
					dp[j][k] = min(dp[j][k], dp[j+1][k]+1)
				}
				if k > 0 {
					dp[j][k] = min(dp[j][k], dp[j][k-1]+1)
				}
				if j+1 < n && k > 0 {
					dp[j][k] = min(dp[j][k], dp[j+1][k-1]+1)
				}
			}
		}
		for j := 0; j < n; j++ {
			for k := n - 1; k >= 0; k-- {
				if j > 0 {
					dp[j][k] = min(dp[j][k], dp[j-1][k]+1)
				}
				if k+1 < n {
					dp[j][k] = min(dp[j][k], dp[j][k+1]+1)
				}
				if j > 0 && k+1 < n {
					dp[j][k] = min(dp[j][k], dp[j-1][k+1]+1)
				}
			}
		}
		for j := n - 1; j >= 0; j-- {
			for k := n - 1; k >= 0; k-- {
				if j+1 < n {
					dp[j][k] = min(dp[j][k], dp[j+1][k]+1)
				}
				if k+1 < n {
					dp[j][k] = min(dp[j][k], dp[j][k+1]+1)
				}
				if j+1 < n && k+1 < n {
					dp[j][k] = min(dp[j][k], dp[j+1][k+1]+1)
				}
			}
		}
		sum := 0
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				sum += dp[j][k]
			}
		}
		fmt.Println(sum)
		Shuffle()
	}
}

func Rand() int {
	X = (A*X + B) % C
	return X
}

func Shuffle() {
	for i := 0; i < n*n; i++ {
		ai := Rand() % n
		aj := Rand() % n
		bi := Rand() % n
		bj := Rand() % n
		if ai == bi && aj == bj {
			continue
		}
		H[ai][aj], H[bi][bj] = H[bi][bj], H[ai][aj]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
