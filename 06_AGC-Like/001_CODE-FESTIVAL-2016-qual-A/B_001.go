package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	cnt := 0
	for i := 1; i <= n; i++ {
		if a[a[i]] == i {
			cnt++
		}
	}
	fmt.Println(cnt / 2)
}
