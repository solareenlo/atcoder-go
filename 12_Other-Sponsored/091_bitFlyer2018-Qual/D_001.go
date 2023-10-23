package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w, n, m int
	fmt.Fscan(in, &h, &w, &n, &m)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}
	H := min(h, n*2+1)
	W := min(w, m*2+1)
	var sum [5000][5000]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s[i][j] == '#' {
				k := H - n + i + 1
				l := W - m + j + 1
				sum[i][j]++
				sum[i][l]--
				sum[k][j]--
				sum[k][l]++
			}
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W-1; j++ {
			sum[i][j+1] += sum[i][j]
		}
	}
	for i := 0; i < H-1; i++ {
		for j := 0; j < W; j++ {
			sum[i+1][j] += sum[i][j]
		}
	}
	ans := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if sum[i][j] != 0 {
				var a int
				if H == n*2+1 && i == n {
					a = h - 2*n
				} else {
					a = 1
				}
				var b int
				if W == m*2+1 && j == m {
					b = w - 2*m
				} else {
					b = 1
				}
				ans += a * b
			}
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
