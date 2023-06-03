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
	A := make([]int, n)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if A[i]+A[j]+A[k] == 1000 {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
}
