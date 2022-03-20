package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		x := -0.5 + (0.002)*float64(i)
		y := -0.5 + (0.002)*float64(i)
		fmt.Println(x, y)
	}
}
