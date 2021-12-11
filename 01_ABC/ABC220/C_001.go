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

	a := make([]int, n)
	sum := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}

	var x int
	fmt.Fscan(in, &x)

	div := x / sum
	rem := x % sum

	cnt := 0
	for rem > 0 {
		rem -= a[cnt]
		cnt++
	}

	if rem == 0 {
		fmt.Println(div*n + cnt + 1)
	} else {
		fmt.Println(div*n + cnt)
	}
}
