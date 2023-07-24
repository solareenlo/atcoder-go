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

	ans := -int(2e9)
	sum := 0
	M := 0
	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(in, &A)
		sum += A
		if ans < sum-M {
			ans = sum - M
		}
		if sum < M {
			M = sum
		}
	}

	fmt.Println(ans)
}
