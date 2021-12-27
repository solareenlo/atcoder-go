package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)

	res := -1
	for i := 0; i < 50001; i++ {
		if math.Floor(float64(i)*1.08) == n {
			res = i
			break
		}
	}

	if res == -1 {
		fmt.Println(":(")
	} else {
		fmt.Println(res)
	}
}
