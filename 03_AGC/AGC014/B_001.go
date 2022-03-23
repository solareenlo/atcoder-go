package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	cnt := make([]int, 100100)
	for i := 1; i <= m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		cnt[a]++
		cnt[b]++
	}

	for i := 1; i <= n; i++ {
		if cnt[i]%2 != 0 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
