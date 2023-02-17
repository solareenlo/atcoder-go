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
	p1 := make([]complex128, n)
	for i := 0; i < n; i++ {
		var x, y float64
		fmt.Fscan(in, &x, &y)
		p1[i] = complex(x, y)
	}
	p2 := make([]complex128, n)
	for i := 0; i < n; i++ {
		var x, y float64
		fmt.Fscan(in, &x, &y)
		p2[i] = complex(x, y)
	}

	c := (p2[1] - p2[0]) / (p1[1] - p1[0])
	r := (p2[0] - c*p1[0]) / (complex(1, 0) - c)

	fmt.Println(real(r), imag(r))
}
