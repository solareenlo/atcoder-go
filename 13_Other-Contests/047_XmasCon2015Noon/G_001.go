package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, a, b int
	var sum int

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b)
		if i > 1 {
			sum += int(math.Abs(float64(a - b)))
		}
		a = b
	}

	fmt.Println(sum)
}
