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

	S := make([]int, 2*N-1)
	a, b := 0, 0
	fmt.Fscan(in, &S[0])
	for i := 1; i < 2*N-1; i++ {
		fmt.Fscan(in, &S[i])
		if S[i] == -1 || S[i] == S[0] {
			continue
		} else if S[i] > S[0] {
			a++
		} else {
			b++
		}
	}

	if max(a, b) >= N {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
