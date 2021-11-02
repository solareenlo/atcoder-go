package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b int
	fmt.Scan(&a, &b)

	res := -100
	for i := 10 * b; i < 10*(b+1); i++ {
		if math.Floor(float64(i)*0.08) == a {
			res = i
			break
		}
	}

	if res == -100 {
		fmt.Println(-1)
	} else {
		fmt.Println(res)
	}
}
