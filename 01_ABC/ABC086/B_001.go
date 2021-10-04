package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	n, _ := strconv.Atoi(a + b)
	sqrt := int(math.Sqrt(float64(n)))
	if sqrt*sqrt == n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
