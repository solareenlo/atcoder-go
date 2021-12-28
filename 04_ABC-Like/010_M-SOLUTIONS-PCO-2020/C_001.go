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

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	for i := k; i < n; i++ {
		if a[i-k] < a[i] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
