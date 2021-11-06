package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var a, b int
	for i := -120; i < 120; i++ {
		for j := -120; j < 120; j++ {
			if i*i*i*i*i-j*j*j*j*j == x {
				a = i
				b = j
				break
			}
		}
	}

	fmt.Println(a, b)
}
