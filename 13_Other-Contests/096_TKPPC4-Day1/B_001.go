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

	var a [200001]int
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := -1
	mx := 0
	for i := 1; i <= N; i++ {
		if K > a[i] && a[i] > mx {
			ans = i
			mx = a[i]
		}
	}
	fmt.Println(ans)
}
