package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	S := make([]string, N)
	for i := range S {
		fmt.Fscan(in, &S[i])
	}

	dh := [4]int{0, 1, 1, 1}
	dw := [4]int{1, 0, 1, -1}
	ok := false
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < 4; k++ {
				c := 0
				for l := 0; l < 6; l++ {
					h := i + dh[k]*l
					w := j + dw[k]*l
					if !(0 <= h && h < N && 0 <= w && w < N) {
						c = 0
						break
					}
					if S[h][w] == '#' {
						c++
					}
				}
				if c >= 4 {
					ok = true
				}
			}
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
