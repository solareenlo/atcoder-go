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

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	ans := 0
	for i := 30; i >= 0; i-- {
		cnt := 0
		for j := 0; j < N; j++ {
			if (A[j]>>i)&1 != 0 {
				cnt++
			}
		}
		if cnt >= 2 {
			ans += (1 << i)
			for j := 0; j < N; j++ {
				if ((A[j] >> i) & 1) == 0 {
					A[j] = 0
				}
			}
		}
	}
	fmt.Println(ans << 1)
}
