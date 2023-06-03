package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, H, W int
	fmt.Fscan(in, &N, &H, &W)
	NIM := 0
	for i := 0; i < N; i++ {
		var A, B int
		fmt.Fscan(in, &A, &B)
		NIM ^= ((A - 1) ^ (B - 1))
	}
	if NIM != 0 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
