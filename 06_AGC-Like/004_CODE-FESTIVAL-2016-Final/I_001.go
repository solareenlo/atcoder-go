package main

import "fmt"

const N = 256

var (
	v   = [N]int{}
	adj = make([][]int, N)
)

func pairs(s string) int {
	ans := 0
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			if s[i] == s[j] {
				ans++
			}
		}
	}
	return ans
}

func dfs(i int) {
	v[i] = 2
	for _, j := range adj[i] {
		if v[j] == 1 {
			dfs(j)
		}
	}
}

func main() {
	var R, C int
	fmt.Scan(&R, &C)
	f := make([]string, R)
	for i := 0; i < R; i++ {
		fmt.Scan(&f[i])
	}

	const mod = 1_000_000_007
	ans := 1
	if R%2 != 0 {
		for i := 0; i < C-1-i; i++ {
			if f[R/2][i] != f[R/2][C-1-i] {
				ans = ans * 2 % mod
				break
			}
		}
	}
	if C%2 != 0 {
		for i := 0; i < R-1-i; i++ {
			if f[i][C/2] != f[R-1-i][C/2] {
				ans = ans * 2 % mod
				break
			}
		}
	}
	cf := []int{12, 12, 6, 4, 0, 0, 1}
	for i := 0; i < R/2; i++ {
		for j := 0; j < C/2; j++ {
			d := pairs(string(f[i][j]) + string(f[R-1-i][j]) + string(f[i][C-1-j]) + string(f[R-1-i][C-1-j]))
			ans = ans * cf[d] % mod
			if d == 0 {
				adj[i] = append(adj[i], R/2+j)
				adj[R/2+j] = append(adj[R/2+j], i)
				v[i] = 1
				v[R/2+j] = 1
			}
		}
	}
	for i := 0; i < R/2+C/2; i++ {
		if v[i] == 1 {
			dfs(i)
			ans = ans * ((mod + 1) / 2) % mod
		}
		if v[i] == 2 {
			ans = ans * 2 % mod
		}
	}
	fmt.Println(ans)
}
