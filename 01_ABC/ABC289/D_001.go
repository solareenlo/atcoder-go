package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var A [10]int
	var dp [100100]int

	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	var M int
	fmt.Fscan(in, &M)
	for i := 0; i < M; i++ {
		var B int
		fmt.Fscan(in, &B)
		dp[B] = -1
	}
	var X int
	fmt.Fscan(in, &X)
	dp[0] = 1
	for i := 0; i < X+1; i++ {
		if dp[i] == 1 {
			for _, Ai := range A {
				if i+Ai <= X && dp[i+Ai] != -1 {
					dp[i+Ai] = 1
				}
			}
		}
	}
	if dp[X] != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
