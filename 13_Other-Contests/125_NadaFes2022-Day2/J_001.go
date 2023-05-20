package main

import "fmt"

func main() {
	var f [3000]int
	f[0] = 1
	f[1] = 2
	for i := 1; i <= 10; i++ {
		t := (1 << i)
		for j := t; j <= 2*t-1; j++ {
			f[j] = 3 - f[j-t]
		}
	}
	vec := make([][]int, 3)
	for i := 0; i < 2048; i++ {
		vec[f[i]] = append(vec[f[i]], i)
	}
	fmt.Printf("%d ", 1024)
	for _, x := range vec[1] {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
	fmt.Printf("%d ", 1024)
	for _, x := range vec[2] {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}
