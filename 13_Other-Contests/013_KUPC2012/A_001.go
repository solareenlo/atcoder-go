package main

import "fmt"

func main() {
	var N, T, E int
	fmt.Scan(&N, &T, &E)
	for i := 0; i < N; i++ {
		var x int
		fmt.Scan(&x)
		for f := T - E; f <= T+E; f++ {
			if f%x == 0 {
				fmt.Println(i + 1)
				return
			}
		}
	}
	fmt.Println(-1)
}
