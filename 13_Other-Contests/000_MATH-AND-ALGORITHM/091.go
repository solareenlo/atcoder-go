package main

import "fmt"

func main() {
	var N, X int
	fmt.Scan(&N, &X)

	ans := 0
	for a := 1; a <= N; a++ {
		for b := a + 1; b <= N; b++ {
			c := X - a - b
			if b < c && c <= N {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
