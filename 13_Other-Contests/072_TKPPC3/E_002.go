package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	b, c := 0.0, 1.0
	for i := 0; i < a; i++ {
		b += (float64(a) + 1) / 2 * c
		c *= (float64(a) - 1 - float64(i)) / float64(a)
	}
	fmt.Println(b)
}
