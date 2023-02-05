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

	var a [210]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := k + 1; i <= n+k; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
}
