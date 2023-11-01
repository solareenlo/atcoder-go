package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		count := 0
		for j := 0; j < a; j++ {
			var b int64
			fmt.Scan(&b)
			if b%2 == 1 {
				count++
			}
		}
		fmt.Println(count)
	}
}
