package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, n)
	sum := 0
	for i := range a {
		fmt.Scan(&a[i])
		sum += a[i]
	}

	threshold := float64(sum) / (4.0 * float64(m))
	t := int(threshold)
	if math.Floor(threshold) == threshold {
		t -= 1
	}

	cnt := 0
	for i := 0; i < n; i++ {
		if t < a[i] {
			cnt++
		}
	}

	if cnt >= m {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
