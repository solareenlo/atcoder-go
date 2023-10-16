package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	a := make([]int, M+1)

	for i := 0; i < N; i++ {
		var b int
		fmt.Fscan(in, &b)
		a[b]++
	}

	ans := -1
	for i := 0; i < M+1; i++ {
		if a[i]*2 > N {
			ans = i
		}
	}

	if ans < 0 {
		fmt.Println("?")
	} else {
		fmt.Println(ans)
	}
}
