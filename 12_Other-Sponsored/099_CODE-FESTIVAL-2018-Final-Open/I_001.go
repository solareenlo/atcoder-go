package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	v := make([][]int, 50)
	K := 0
	for N > 0 {
		N--
		var L, R int
		fmt.Fscan(in, &L, &R)
		v[L] = append(v[L], R)
		K += 1 << L
	}
	L, R := 0, K
	for R-L > 1 {
		K = (R + L) / 2
		N = 0
		w := make([][]int, len(v))
		for i := range v {
			w[i] = make([]int, len(v[i]))
			copy(w[i], v[i])
		}
		for i := 0; i < 49; i, K = i+1, K>>1 {
			if (K & 1) != 0 {
				w[i] = append(w[i], 0)
			}
			sort.Slice(w[i], func(a, b int) bool {
				return w[i][a] > w[i][b]
			})
			if (K & 1) != 0 {
				N += w[i][0]
			}
			for j := K & 1; j+1 < len(w[i]); j += 2 {
				w[i+1] = append(w[i+1], w[i][j]+w[i][j+1])
			}
		}
		K = (R + L) / 2
		if N < M {
			L = K
		} else {
			R = K
		}
	}
	fmt.Println(R)
}
