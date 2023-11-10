package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200005

	var a, b [N]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sum := 0
	for i := 1; i <= n; i++ {
		b[i] = i - 1e6
		sum += a[i] * b[i]
	}
	for i, s := n, 0; i >= 1; i-- {
		s += a[i]
		if (s == -1 && sum > 0) || (s == 1 && sum < 0) {
			for j := i; j <= n; j++ {
				b[j] += abs(sum)
			}
			sum = 0
			break
		}
	}
	if sum != 0 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", b[i])
	}
	fmt.Println()
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
