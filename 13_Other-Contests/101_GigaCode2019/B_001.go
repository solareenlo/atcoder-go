package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, X, Y, Z int
	fmt.Fscan(in, &N, &X, &Y, &Z)
	cnt := 0
	for i := 0; i < N; i++ {
		var A, B int
		fmt.Fscan(in, &A, &B)
		if A >= X && B >= Y && A+B >= Z {
			cnt++
		}
	}
	fmt.Println(cnt)
}
