package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b float64
	fmt.Fscan(in, &a, &b)

	x := float64(int(math.Pow(a/(b*2), 2.0/3.0)) - 1)
	t := x*b + a/math.Sqrt(x+1)
	s := (x+1)*b + a/math.Sqrt(x+2)
	fmt.Println(math.Min(math.Min(s, t), a))
}
