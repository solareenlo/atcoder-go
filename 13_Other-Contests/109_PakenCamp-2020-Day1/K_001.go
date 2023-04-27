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

	total := 0
	var a [305]int
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i])
		total += a[i]
	}

	var ans [305]int
	for mul := 1; mul <= 300*300; mul++ {
		if total%mul != 0 {
			continue
		}
		cnt := 0
		sum := 0
		for i := 0; i < N; i++ {
			sum = (sum + a[i]) % mul
			if sum == 0 {
				cnt++
			}
		}
		ans[cnt] = mul
	}

	for i := N - 1; i >= 1; i-- {
		ans[i] = max(ans[i], ans[i+1])
	}
	for i := 1; i <= N; i++ {
		fmt.Println(ans[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
