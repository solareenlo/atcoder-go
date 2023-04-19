package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MX = 1000005

	var n int
	fmt.Fscan(in, &n)

	var p [MX]int
	var np [MX]bool
	for i := 3; i < MX; i += 2 {
		if !np[i] {
			p[i] = p[i-2] + 1
			for j := i; j < MX; j += i {
				np[j] = true
			}
		}
	}
	p[2] = 1
	o := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		o ^= p[x]
	}
	if o != 0 {
		fmt.Println("An")
	} else {
		fmt.Println("Ai")
	}
}
