package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	var A [100001]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i])
	}
	ans := 0
	j := 1
	for i := 0; i < n-1; i++ {
		for A[j]-A[i] <= k && j < n {
			j++
		}
		ans += j - i - 1
	}
	fmt.Println(ans)
}
