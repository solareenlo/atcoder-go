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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	X := make(map[int]bool)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				X[a[i]*a[j]*a[k]] = true
			}
		}
	}
	fmt.Println(len(X))
}
