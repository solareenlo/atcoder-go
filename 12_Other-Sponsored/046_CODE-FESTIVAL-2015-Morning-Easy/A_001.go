package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := n; ; i++ {
		a := math.Sqrt(float64(i))
		if float64(int(math.Sqrt(float64(i)))) == a {
			fmt.Println(i - n)
			return
		}
	}
}
