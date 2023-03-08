package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000009

	var H, W int
	fmt.Fscan(in, &H, &W)
	m := make(map[int]int)
	ind := 0
	for i := 0; i < H; i++ {
		var c string
		fmt.Fscan(in, &c)
		for j := 0; j < W; j++ {
			if c[j] == 'W' {
				m[i*W+j] = ind
				ind++
			}
		}
	}
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	G := make([][]int, ind)
	for i := range G {
		G[i] = make([]int, ind)
	}
	for x, y := range m {
		for k := 0; k < 4; k++ {
			s := x / W
			t := x % W
			for {
				s += dx[k]
				t += dy[k]
				if min(s, t) == -1 || s == H || t == W {
					break
				}
				if _, ok := m[s*W+t]; ok {
					G[y][m[s*W+t]] = 1
					break
				}
			}
		}
	}
	ans := Kirchhoffs_theorem(G, mod)
	fmt.Println(ans)
}

// 行列木定理
func Kirchhoffs_theorem(G [][]int, MOD int) int {
	N := len(G)
	H := make([][]int, N-1)
	for i := range H {
		H[i] = make([]int, N-1)
	}
	for i := 0; i < N-1; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			H[i][i] = (H[i][i] + G[i][j]) % MOD
			if j != N-1 {
				H[i][j] = (MOD - G[i][j]) % MOD
			}
		}
	}
	return Determinant_Matrix(H, MOD)
}

func Determinant_Matrix(G [][]int, MOD int) int {
	N := len(G)
	ans := 1
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			G[j][i] = (G[j][i]%MOD + MOD) % MOD
			if G[j][i] != 0 {
				if j != i {
					for k := i; k < N; k++ {
						G[i][k], G[j][k] = G[j][k], G[i][k]
					}
					ans *= -1
				}
				break
			}
		}
		if G[i][i] == 0 {
			return 0
		}
		ans = (ans * G[i][i]) % MOD
		R := rev(G[i][i], MOD)
		for j := i + 1; j < N; j++ {
			D := (R * G[j][i]) % MOD
			for k := i; k < N; k++ {
				G[j][k] = (G[j][k] - (D*G[i][k])%MOD + MOD) % MOD
			}
		}
	}
	return (ans + MOD) % MOD
}

func rev(a int, MOD int) int {
	a %= MOD
	ans := 1
	H := MOD - 2
	for H != 0 {
		if (H & 1) != 0 {
			ans = (ans * a) % MOD
		}
		a = (a * a) % MOD
		H >>= 1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
