package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Float64s(a[1:])
	a = append(a, a[1])
	a = append(a, a[2])
	s := -9999999999999999.0
	t := 9999999999999999.0
	for i := 1; i <= n; i++ {
		x := (a[i] + a[i+1] + a[i+2]) / a[i] / a[i+1] / a[i+2]
		s = math.Max(s, x)
		t = math.Min(t, x)
	}
	fmt.Println(t)
	fmt.Println(s)
}
