package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100010

	var arr [N]int

	var n, q int
	fmt.Fscan(in, &n, &q)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &arr[i])
		arr[i] += arr[i-1]
	}

	for i := 1; i <= q; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		fmt.Printf("%d\n", arr[y]-arr[x-1])
	}
}
