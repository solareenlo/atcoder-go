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
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	for i := 2; i <= n; i++ {
		if abs(a[i-1]-a[i]) > k && abs(b[i-1]-a[i]) > k {
			a[i] = -1
		}
		if abs(a[i-1]-b[i]) > k && abs(b[i-1]-b[i]) > k {
			b[i] = -1
		}
		if a[i] == -1 && b[i] == -1 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
