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

	res := 0.0
	for i := 0; i < n; i++ {
		var p, q float64
		fmt.Fscan(in, &p, &q)
		res += q / p
	}

	fmt.Println(res)
}
