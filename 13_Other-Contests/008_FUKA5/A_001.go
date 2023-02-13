package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	for {
		var n, k int
		fmt.Fscan(in, &n, &k)
		if n == 0 && k == 0 {
			break
		}

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		sort.Slice(a, func(i, j int) bool {
			return a[i] < a[j]
		})

		cnt := 0
		for i := 0; i < k; i++ {
			cnt += a[i]
		}
		fmt.Println(cnt)
	}
}
