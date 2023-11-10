package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	for n > 0 {
		n--
		var x int
		fmt.Fscan(in, &x)
		if x == 2 && k < x {
			fmt.Println("No")
			return
		}
		k -= (x + 1)
	}
	fmt.Println("Yes")
}
