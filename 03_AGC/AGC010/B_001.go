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

	a := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}

	p := (n * (n + 1) / 2)
	if sum%p != 0 {
		fmt.Println("NO")
		return
	}

	b := make([]int, n)
	k := sum / p
	for i := 0; i < n; i++ {
		b[i] = a[(i+1)%n] - a[i] - k
	}

	for i := 0; i < n; i++ {
		if b[i] > 0 || (-b[i])%n != 0 {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
