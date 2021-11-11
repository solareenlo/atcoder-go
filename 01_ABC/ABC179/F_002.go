package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, Q int
	fmt.Fscan(in, &N, &Q)

	a := make([]int, N+1)
	b := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		a[i] = N
		b[i] = N
	}

	black := (N - 2) * (N - 2)
	H, W := N, N
	for ; Q > 0; Q-- {
		var t, x int
		fmt.Fscan(in, &t, &x)
		if t == 1 {
			if x < W {
				for i := x; i < W; i++ {
					b[i] = H
				}
				W = x
			}
			black -= b[x] - 2

		} else {
			if x < H {
				for i := x; i < H; i++ {
					a[i] = W
				}
				H = x
			}
			black -= a[x] - 2
		}
	}

	fmt.Println(black)
}
