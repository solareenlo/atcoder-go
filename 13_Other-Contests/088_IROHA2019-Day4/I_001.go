package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = i
	}
	sort.Slice(ans, func(u, v int) bool {
		fmt.Printf("1 %d %d\n", ans[u]+1, ans[v]+1)
		out.Flush()
		fmt.Printf("2 %d %d\n", ans[u]+1, ans[v]+1)
		out.Flush()
		var buf string
		fmt.Scan(&buf)
		fmt.Printf("1 %d %d\n", ans[u]+1, ans[v]+1)
		out.Flush()
		return buf[0] == 'Y'
	})
	fmt.Printf("0 %d 0\n", n)
	out.Flush()
	for i := 0; i < n; i++ {
		fmt.Printf("%d\n", ans[i]+1)
		out.Flush()
	}
}
