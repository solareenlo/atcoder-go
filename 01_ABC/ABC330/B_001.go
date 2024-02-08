package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200005

	var n, l, r int
	fmt.Fscan(in, &n, &l, &r)
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if a < l {
			fmt.Printf("%d ", l)
		} else if a > r {
			fmt.Printf("%d ", r)
		} else {
			fmt.Printf("%d ", a)
		}
	}
}
