package main

import "fmt"

var (
	S   = [12][12]bool{}
	vis = [12][12]bool{}
)

func DFS(r, c int) int {
	res := 0
	if vis[r][c] {
		return 0
	}
	vis[r][c] = true
	if S[r+1][c] {
		res += DFS(r+1, c)
	}
	if S[r-1][c] {
		res += DFS(r-1, c)
	}
	if S[r][c+1] {
		res += DFS(r, c+1)
	}
	if S[r][c-1] {
		res += DFS(r, c-1)
	}
	return res + 1
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	sum := 1
	for i := 1; i <= N; i++ {
		var s string
		fmt.Scan(&s)
		for j := 1; j <= M; j++ {
			if s[j-1] == '.' {
				S[i][j] = true
			}
			if S[i][j] {
				sum++
			}
		}
	}

	ans := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			for r := 1; r <= N; r++ {
				for c := 1; c <= M; c++ {
					vis[r][c] = false
				}
			}
			if !S[i][j] {
				if DFS(i, j) == sum {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
