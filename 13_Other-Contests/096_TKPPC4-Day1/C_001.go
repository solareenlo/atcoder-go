package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	var X string
	fmt.Fscan(in, &N, &X)

	s := len(X)
	for i := 2; i <= 10; i++ {
		t := 0
		for k := 0; k < s; k++ {
			a := int(X[k] - '0')
			for j := 0; j < s-k-1; j++ {
				a *= i
			}
			t += a
		}
		if t == N {
			fmt.Println(i)
			break
		}
	}
	return
}
