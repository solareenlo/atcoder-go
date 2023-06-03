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

	var d, n int
	fmt.Fscan(in, &d, &n)
	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		arr[x]++
		arr[y+1]--
	}
	for i := 1; i <= d; i++ {
		arr[i] += arr[i-1]
		fmt.Println(arr[i])
	}
}
