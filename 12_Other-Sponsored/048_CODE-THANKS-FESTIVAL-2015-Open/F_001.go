package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	var sum [100005]int
	for i := 0; i < N-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		sum[a]++
		sum[b]++
	}
	if sum[1] == 1 {
		fmt.Println("A")
	} else {
		if N%2 == 0 {
			fmt.Println("A")
		} else {
			fmt.Println("B")
		}
	}
}
