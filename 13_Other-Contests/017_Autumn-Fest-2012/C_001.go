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

	a := make([]float64, 1005)
	sum := 0.0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}
	sum /= float64(n)
	sum *= float64(n - 1)
	fmt.Println(sum)
}
