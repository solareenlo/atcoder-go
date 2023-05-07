package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, P, Q int
	fmt.Fscan(in, &N, &P, &Q)
	A := make([]int, N)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}

	count := 0
	for i := 0; i < N-4; i++ {
		a := A[i]
		for j := (i + 1); j < (N - 3); j++ {
			b := a * A[j] % P
			for k := (j + 1); k < (N - 2); k++ {
				c := b * A[k] % P
				for l := (k + 1); l < (N - 1); l++ {
					d := c * A[l] % P
					for m := (l + 1); m < N; m++ {
						e := d * A[m] % P
						if e == Q {
							count++
						}
					}
				}
			}
		}
	}
	fmt.Println(count)
}
