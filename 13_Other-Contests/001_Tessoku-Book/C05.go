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
	var A [10]int
	for i := 0; i < 10; i++ {
		A[i] = 4
	}
	n--
	i := 0
	for n > 0 {
		if (n & 1) != 0 {
			A[i] = 7
		}
		i++
		n >>= 1
	}
	for i := 9; i >= 0; i-- {
		fmt.Print(A[i])
	}
}
