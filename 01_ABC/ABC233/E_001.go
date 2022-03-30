package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a string
	fmt.Fscan(in, &a)

	ten := make([]int, 1000000)
	n := len(a)
	for i, j := n-1, 0; i >= 0; i, j = i-1, j+1 {
		ten[i] = ten[i+1] + int(a[j]-'0')
	}

	for i := 0; i < n-1; i++ {
		ten[i+1] += ten[i] / 10
		ten[i] %= 10
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Fprint(out, ten[i])
	}
	fmt.Fprintln(out)
}
