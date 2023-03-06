package main

import (
	"fmt"
)

func main() {
	era := make([]bool, 1001)
	for i := 2; i <= 1000; i++ {
		era[i] = true
	}

	ans := 1
	for i := 2; i <= 1000; i++ {
		if era[i] {
			for j := i * 2; j <= 1000; j += i {
				era[j] = false
			}
			for j := i; j <= 1000; j *= i {
				fmt.Println("?", j)
				var c string
				fmt.Scan(&c)
				if c[0] == 'Y' {
					ans *= i
				}
			}
		}
	}
	fmt.Println("!", ans)
}
