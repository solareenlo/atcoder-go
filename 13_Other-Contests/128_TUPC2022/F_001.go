package main

import (
	"bufio"
	"fmt"
	"os"
)

var N int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N)
	var isp [1 << 17]int
	Ps := make([]int, 0)
	for i := 2; i < 1<<17; i++ {
		if isp[i] == 0 {
			Ps = append(Ps, i)
			for j := i; j < 1<<17; j += i {
				isp[j] = len(Ps)
			}
		}
	}
	cnt := make([]int, len(Ps))
	H := make([]int, N)
	W := make([]int, N)
	A := make([][]int, N)
	sz := 0
	C := make([]int, 0)
	for i := 0; i < N; i++ {
		var M int
		fmt.Fscan(in, &M)
		H[i] = M
		A[i] = make([]int, M)
		for j := 0; j < M; j++ {
			var c int
			fmt.Fscan(in, &c)
			C = append(C, c)
			A[i][j] = sz
			sz++
			W[j]++
		}
	}
	B := rot(A)
	for i := 0; i < N; i++ {
		if H[i] != W[i] {
			B = rot(B)
			break
		}
	}
	P := make([]int, len(C))
	for i := 0; i < N; i++ {
		for j := 0; j < len(A[i]); j++ {
			P[A[i][j]] = B[i][j]
		}
	}
	var vis [1 << 17]bool
	for i := 0; i < len(P); i++ {
		if !vis[i] {
			Q := make([]int, 0)
			u := i
			for !vis[u] {
				vis[u] = true
				Q = append(Q, C[u])
				u = P[u]
			}
			g := len(Q)
			for d := 1; d < len(Q); d++ {
				if len(Q)%d == 0 {
					ok := true
					for i := 0; i < len(Q); i++ {
						if Q[i] != Q[(i+d)%len(Q)] {
							ok = false
							break
						}
					}
					if ok {
						g = d
						break
					}
				}
			}
			for g > 1 {
				pi := isp[g] - 1
				p := Ps[pi]
				c := 0
				for g%p == 0 {
					g /= p
					c++
				}
				cnt[pi] = max(cnt[pi], c)
			}
		}
	}
	const mod = 998244353
	ans := 1
	for i := 0; i < len(Ps); i++ {
		for j := 0; j < cnt[i]; j++ {
			ans = ans * Ps[i] % mod
		}
	}
	for i := 0; i < N; i++ {
		if H[i] != W[i] {
			ans = ans * 2 % mod
			break
		}
	}
	fmt.Println(ans)
}

func rot(A [][]int) [][]int {
	B := make([][]int, N)
	for i := N - 1; i >= 0; i-- {
		for j := 0; j < len(A[i]); j++ {
			B[j] = append(B[j], A[i][j])
		}
	}
	return B
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
