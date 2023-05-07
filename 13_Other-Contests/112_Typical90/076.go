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

	var v [123456]int
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v[i])
		sum += v[i]
	}

	now := 0
	nowidx := 0
	for i := 0; i < n; i++ {
		for now*10 < sum {
			now += v[(nowidx)%n]
			nowidx++
		}
		if now*10 == sum {
			fmt.Println("Yes")
			return
		} else {
			now -= v[i]
		}
	}
	fmt.Println("No")
}
