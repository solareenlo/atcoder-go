package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const maxn = 50000000000000
	var n, k int
	fmt.Scan(&n, &k)
	for i := 1; i < n; i++ {
		fmt.Printf("%d %d ", maxn-i, maxn+i)
	}
	for i := 1; i < k-(n-1)*2; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println(maxn*2 - (k-(n-1)*2)*(k-(n-1)*2-1)/2)
	fmt.Println("YES")
	for i := 1; i < n; i++ {
		fmt.Printf("%d %d ", i, i)
	}
	for i := n*2 - 1; i <= k; i++ {
		fmt.Printf("%d ", n)
	}
}
