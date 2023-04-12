package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	s := 0.0
	for i := 1; i < n; i++ {
		var t float64
		fmt.Fscan(in, &t)
		s += t
	}
	fmt.Println(math.Ceil(s / float64(m)))
}
