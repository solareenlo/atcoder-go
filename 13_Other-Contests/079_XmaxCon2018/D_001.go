package main

import "fmt"

var N, K int
var ans []int

func main() {
	fmt.Scan(&N, &K)

	for i := 1; i < K*K+1; i++ {
		chk(i)
	}

	rest := N - len(ans)
	for i := 0; i < rest; i++ {
		for j := 0; j < K; j++ {
			if j == K-1 {
				fmt.Printf("1\n")
			} else {
				fmt.Printf("1 ")
			}
		}
	}
	const M = 10000
	idx := 3
	for _, big := range ans {
		ar := make([]int, 0)
		for i := 0; i < big; i++ {
			ar = append(ar, M+idx)
		}
		for i := 0; i < K-big; i++ {
			ar = append(ar, idx)
		}
		idx++
		for i := 0; i < K; i++ {
			if i == K-1 {
				fmt.Printf("%d\n", ar[i])
			} else {
				fmt.Printf("%d ", ar[i])
			}
		}
	}
}

func chk(L int) bool {
	L2 := K*K - L
	for n := 2; n <= N; n++ {
		for i := 1; i <= K; i++ {
			lrg := i
			tmp := make([]int, 0)
			tmp = append(tmp, lrg)
			for j := 1; j < n && lrg > 0; j++ {
				sml2 := L2 / lrg
				lrg = min(lrg, K-sml2)
				tmp = append(tmp, lrg)
			}
			if lrg > 0 && (K-lrg)*i >= L {
				ans = tmp
				return true
			}
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
