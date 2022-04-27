package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	S := [32][32]int{}
	for i := 1; i <= N; i++ {
		var s string
		fmt.Scan(&s)
		for j := 1; j <= M; j++ {
			if s[j-1] == '#' {
				S[i][j] = 1
			}
		}
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			ans := 0
			for ii := -1; ii <= 1; ii++ {
				for jj := -1; jj <= 1; jj++ {
					ans += S[i+ii][j+jj]
				}
			}
			fmt.Print(ans)
		}
		fmt.Println()
	}
}
