package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var x int
	var a [15]int
	for i := 1; i <= 5; i++ {
		fmt.Fscan(in, &x)
		a[x]++
	}

	f := false
	for i := 1; i < 15; i++ {
		for j := 1; j < 15; j++ {
			if a[i] == 2 && a[j] == 3 {
				f = true
			}
		}
	}

	if f {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
