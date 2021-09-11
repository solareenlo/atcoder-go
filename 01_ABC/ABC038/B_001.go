package main

import "fmt"

func main() {
	a, b := make([]int, 2), make([]int, 2)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for i := range b {
		fmt.Scan(&b[i])
	}

	res := "NO"
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if a[i] == b[j] {
				res = "YES"
			}
		}
	}
	fmt.Println(res)
}
