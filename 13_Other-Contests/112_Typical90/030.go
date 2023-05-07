package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)

	soin := make([]int, 1<<25)
	for n := 2; n < N+1; n++ {
		if soin[n] != 0 {
			continue
		}
		for i := n; i <= N; i += n {
			soin[i]++
		}
	}

	ans := 0
	for n := 2; n < N+1; n++ {
		if soin[n] >= K {
			ans++
		}
	}
	fmt.Println(ans)
}
