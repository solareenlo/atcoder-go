package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const B = 340

	var N, M, Q int
	fmt.Fscan(in, &N, &M, &Q)
	BN := (N + B - 1) / B
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
		A[i] %= M
	}
	var cnt [B][100000]int
	for i := 0; i < BN; i++ {
		a := i * B
		b := min(a+B, N)
		for j := a; j < b; j++ {
			cnt[i][A[j]]++
		}
	}
	var offset [B]int
	for i := 0; i < Q; i++ {
		var l, r, d int
		fmt.Fscan(in, &l, &r, &d)
		l--
		d %= M
		ans := 0
		lo := l/B - 1
		hi := r/B + 1
		for j := lo; j < hi; j++ {
			a := j * B
			b := a + B
			if b <= l || r <= a {
				continue
			}
			if l <= a && b <= r {
				offset[j] -= d
				if offset[j] < 0 {
					offset[j] += M
				}
				ans += cnt[j][offset[j]]
			} else {
				x := max(a, l)
				y := min(b, r)
				for k := x; k < y; k++ {
					cnt[j][A[k]]--
					A[k] += d
					if A[k] >= M {
						A[k] -= M
					}
					cnt[j][A[k]]++
					if A[k] == offset[j] {
						ans++
					}
				}
			}
		}
		fmt.Println(ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
