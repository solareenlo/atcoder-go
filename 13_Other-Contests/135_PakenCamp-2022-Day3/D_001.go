package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	sum := 0
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
		sum += A[i]
	}
	K := 0
	for sum%2 == 0 {
		K++
		sum >>= 1
	}
	for i := 0; i < N; i++ {
		if A[i]%sum != 0 {
			fmt.Println("No")
			return
		}
		A[i] /= sum
	}
	ops := make([]pair, 0)
	for k := 0; k < K; k++ {
		Ai := make([]pair, 0)
		for i := 0; i < N; i++ {
			if ((A[i] >> k) & 1) != 0 {
				Ai = append(Ai, pair{A[i], i})
			}
		}
		sortPair(Ai)
		for i := 0; i < len(Ai); i += 2 {
			x := Ai[i].y
			y := Ai[i+1].y
			ops = append(ops, pair{x, y})
			A[y] -= A[x]
			A[x] *= 2
		}
	}
	fmt.Println("Yes")
	fmt.Println(len(ops))
	for _, xy := range ops {
		fmt.Println(xy.x+1, xy.y+1)
	}
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}
