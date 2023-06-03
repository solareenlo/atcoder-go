package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	S := make([]pair, 0)
	for i := 0; i < n; i++ {
		d := -2
		for len(S) != 0 {
			if S[len(S)-1].x > a[i] {
				d = S[len(S)-1].y
				break
			} else {
				S = S[:len(S)-1]
			}
		}
		S = append(S, pair{a[i], i})
		if i == n-1 {
			fmt.Println(d + 1)
		} else {
			fmt.Printf("%d ", d+1)
		}
	}
}
